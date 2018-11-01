package lnk

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
)

// LinkInfoSection represents the LinkInfo structure. Section 2.3 of [MS-SHLLINK].
// It appears right after LinkTargetIDList if it's in the linkFlags.
type LinkInfoSection struct {
	// LinkInfo header section start.

	// Size of the LinkInfo structure. Includes these four bytes.
	LinkInfoSize uint32

	// Size of LinkInfo header section.
	// If == 0x1c => Offsets to optional fields are not specified.
	// If >= 0x24 => Offsets to optional fields are specified.
	// Header section include LinkInfoSize and some of the following fields.
	LinkInfoHeaderSize uint32

	// Offsets are from the start of LinkInfo structure == start of the io.Reader.
	// If an offset is zero, then field does not exist.

	// Flags that specify whether the VolumeID, LocalBasePath, LocalBasePathUnicode,
	// and CommonNetworkRelativeLink fields are present in this structure.
	// See linkInfoFlags
	LinkInfoFlags uint32

	// LinkInfoFlagsStr contains the flags in string format.
	LinkInfoFlagsStr []string

	// Offset of VolumeID if VolumeIDAndLocalBasePath is set.
	VolumeIDOffset uint32

	// Offset of LocalBasePath if VolumeIDAndLocalBasePath is set.
	LocalBasePathOffset uint32

	// Offset of CommonNetworkRelativeLink if CommonNetworkRelativeLinkAndPathSuffix is set.
	CommonNetworkRelativeLinkOffset uint32

	// Offset of CommonPathSuffix.
	CommonPathSuffixOffset uint32

	// Offset of optional LocalBasePathUnicode and present if VolumeIDAndLocalBasePath is set
	// and LinkInfoHeaderSize >= 0x24.
	LocalBasePathOffsetUnicode uint32 // Optional

	// Offset of CommonPathSuffixUnicode and present if LinkInfoHeaderSize >= 0x24.
	CommonPathSuffixOffsetUnicode uint32 // Optional

	// LinkInfo header section end. At least I think.

	// VolumeID present if VolumeIDAndLocalBasePath is set.
	VolID VolumeID

	// Null-terminated string present if VolumeIDAndLocalBasePath is set.
	// Combine with CommonPathSuffix to get the full path to target.
	LocalBasePath string // Optional

	// Optional CommonNetworkRelativeLink, contains information about network
	// location of the target.
	NetworkRelativeLink CommonNetworkRelativeLink

	// Null-terminated string. Combine with LocalBasePath to get full path to target.
	CommonPathSuffix string // Optional

	// Null-terminated Unicode string to base path.
	// Present only VolumeIDAndLocalBasePath is set and LinkInfoHeaderSize >= 0x24.
	LocalBasePathUnicode string // Optional

	// Null-terminated Unicode string to common path.
	// Present only VolumeIDAndLocalBasePath is set and LinkInfoHeaderSize >= 0x24.
	CommonPathSuffixUnicode string // Optional
}

// linkInfoFlags defines the LinkInfoFlags. Only the first two bits are used for now.
var linkInfoFlags = []string{
	// If 1, VolumeIDOffset and LocalBasePathOffset point to respective fields.
	// If LinkInfoHeaderSize >= 0x24 and LocalBasePathOffsetUnicode is populated.
	"VolumeIDAndLocalBasePath", // Bit 0

	// If 1, CommonNetworkRelativeLinkOffset field is populated.
	// If 0, offset is zero.
	"CommonNetworkRelativeLinkAndPathSuffix", // Bit 1
}

// LinkInfo reads the io.Reader and returns a populated LinkInfoSection.
func LinkInfo(r io.Reader) (info LinkInfoSection, err error) {
	// Read size.
	err = binary.Read(r, binary.LittleEndian, &info.LinkInfoSize)
	if err != nil {
		return info, fmt.Errorf("lnk.LinkInfo: read LinkInfoSize - %s", err.Error())
	}

	// This time we have offsets, so it's better to read the complete LinkInfo
	// and store it in a []byte for easy access to offset.
	// We must read (LinkInfoSize-4) bytes from io.Reader and append it to
	// LinkInfoSize to get the complete []byte.

	fmt.Printf("LinkInfoSize: %x\n", info.LinkInfoSize)

	tempInfo := make([]byte, info.LinkInfoSize)

	// Does order matter here? It seems like order does not matter when reading
	// byte slices or arrays. Let's go with BigEndian and see what happens. Doesn't matter.
	err = binary.Read(r, binary.LittleEndian, tempInfo)
	if err != nil {
		return info, fmt.Errorf("lnk.LinkInfo: read LinkInfo - %s", err.Error())
	}

	linkInfo := uint32Byte(info.LinkInfoSize)
	linkInfo = append(linkInfo, tempInfo...)

	fmt.Printf("linkInfo:\n%s\n", hex.Dump(linkInfo))

	// Convert it into a reader to read the header, then we can use the []byte
	// for reading offsets.
	buf := bytes.NewReader(linkInfo)

	// Read the first four size bytes again. This is to move the reader forward.
	err = binary.Read(buf, binary.LittleEndian, &info.LinkInfoSize)
	if err != nil {
		return info, fmt.Errorf("lnk.LinkInfo: read LinkInfoSize-2 - %s", err.Error())
	}

	// Read LinkInfoHeaderSize.
	err = binary.Read(buf, binary.LittleEndian, &info.LinkInfoHeaderSize)
	if err != nil {
		return info, fmt.Errorf("lnk.LinkInfo: read LinkInfoHeaderSize - %s", err.Error())
	}

	// If 0x1C no optional fields.
	// If >= 0x24 offset to optional fields are here.
	optionalHeaderFields := false
	if info.LinkInfoHeaderSize == 0x1c {
		optionalHeaderFields = false
	}
	if info.LinkInfoHeaderSize >= 0x24 {
		optionalHeaderFields = true
	}

	fmt.Printf("LinkInfoHeaderSize is %x so setting optionalHeaderFields to %v.\n", info.LinkInfoHeaderSize, optionalHeaderFields)

	// Read LinkInfoFlags.
	err = binary.Read(buf, binary.LittleEndian, &info.LinkInfoFlags)
	if err != nil {
		return info, fmt.Errorf("lnk.LinkInfo: read LinkInfoFlags - %s", err.Error())
	}

	fmt.Println("LinkInfoFlags", info.LinkInfoFlags)
	// Set flags.
	for bitIndex := 0; bitIndex < 2; bitIndex++ {
		if bitMaskuint32(info.LinkInfoFlags, bitIndex) {
			info.LinkInfoFlagsStr = append(info.LinkInfoFlagsStr, linkInfoFlags[bitIndex])
		}
	}

	fmt.Println("LinkInfoFlagsStr", info.LinkInfoFlagsStr)

	// If VolumeIDAndLocalBasePath is set then VolumeIDOffset is set.
	if bitMaskuint32(info.LinkInfoFlags, 0) {
		// Read VolumeIDOffset.
		err = binary.Read(buf, binary.LittleEndian, &info.VolumeIDOffset)
		if err != nil {
			return info, fmt.Errorf("lnk.LinkInfo: read VolumeIDOffset - %s", err.Error())
		}
	}

	return info, err
}
