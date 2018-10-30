package lnk

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

const (
	headerSize = 0x4C
	// This is how it appears in the file when we read an array.
	// Endian-ness does not make a difference.
	classID = "0114020000000000c000000000000046"
)

// From the docs:
// "Multi-byte data values in the Shell Link Binary File Format are stored in little-endian format."

// ReadFile reads an lnk file and populates the header, data, and size fields.

// Header parses the first 0x4c bytes of the slice and returns a ShellLinkHeader.
func Header(b []byte) (head ShellLinkHeader, err error) {
	if len(b) < headerSize {
		return head, fmt.Errorf("lnk.header: small input - got %d, want %d", len(b), 0x4c)
	}

	// Make an io.Reader from header bytes. Makes life easier.
	buf := bytes.NewReader(b[:headerSize])

	// Check the first four bytes against 0x4c.
	var magic uint32
	err = binary.Read(buf, binary.LittleEndian, &magic)
	if err != nil {
		return head, fmt.Errorf("lnk.header: error reading magic string - %s", err.Error())
	}
	if magic != headerSize {
		return head, fmt.Errorf("lnk.header: invalid magic string- got %x, want %s", magic, "0x4C")
	}
	head.Header = magic

	// Next 16 bytes should be 00021401-0000-0000-C000-000000000046.
	// Read two uint64 and compare.
	var clsID [16]byte
	err = binary.Read(buf, binary.LittleEndian, &clsID)
	if err != nil {
		return head, fmt.Errorf("lnk.header: error reading LinkCLSID - %s", err.Error())
	}
	hexClsID := hex.EncodeToString(clsID[:])
	if hexClsID != classID {
		return head,
			fmt.Errorf("lnk.header: invalid magic string- got %s, want %s", hexClsID, classID)
	}
	head.LinkCLSID = clsID

	// Parse LinkFlags.
	// Convert the next uint32 to bites, go over the bits and add the flags.
	var flags []string
	// We will lose preceeding zeros by using uint32 but we do not care about them.
	var lf uint32
	err = binary.Read(buf, binary.LittleEndian, &lf)
	if err != nil {
		return head, fmt.Errorf("lnk.header: reading LinkFlags - %s", err.Error())
	}
	flags = matchFlag(lf, linkFlags)
	head.LinkFlags = flags
	fmt.Println(flags)

	// Parse FileAttributes.
	var attribs uint32
	// Same as before, read BigEndian.
	err = binary.Read(buf, binary.LittleEndian, &attribs)
	if err != nil {
		return head, fmt.Errorf("lnk.header: reading FileAttributes - %s", err.Error())
	}
	fmt.Println(matchFlag(attribs, fileAttributesFlags))

	// Convert timestamps from Windows Filetime to time.Time.
	var crTime, wrTime, acTime [8]byte
	err = binary.Read(buf, binary.BigEndian, &crTime)
	if err != nil {
		return head, fmt.Errorf("lnk.header: reading CreationTime - %s", err.Error())
	}
	head.CreationTime = toTime(crTime)

	err = binary.Read(buf, binary.BigEndian, &wrTime)
	if err != nil {
		return head, fmt.Errorf("lnk.header: reading WriteTime - %s", err.Error())
	}
	head.WriteTime = toTime(wrTime)

	err = binary.Read(buf, binary.BigEndian, &acTime)
	if err != nil {
		return head, fmt.Errorf("lnk.header: reading AccessTime - %s", err.Error())
	}
	head.AccessTime = toTime(acTime)

	fmt.Println(head.CreationTime)
	fmt.Println(head.WriteTime)
	fmt.Println(head.AccessTime)

	// var size uint32
	var size [4]byte
	err = binary.Read(buf, binary.LittleEndian, &size)
	if err != nil {
		return head, fmt.Errorf("lnk.header: reading target file size - %s", err.Error())
	}

	fmt.Printf("%b", size)

	return head, err
}

// // timeFromWinToUnix converts a Windows style timestamp to Unix.
// func timeFromWinToUnix(t uint64) t {

// }
