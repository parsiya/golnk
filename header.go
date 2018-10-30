package lnk

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"syscall"
	"time"
)

const (
	headerSize = 0x4C
	// This is how it appears in the file when we read an array.
	// Endian-ness does not make a difference.
	classID = "0114020000000000c000000000000046"
)

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
	var lf uint32 // We will lose zeros by using uint32 but we do not care about them.
	err = binary.Read(buf, binary.LittleEndian, &lf)
	if err != nil {
		return head, fmt.Errorf("lnk.header: error reading LinkFlags - %s", err.Error())
	}
	// Convert to bits. TODO: Better way than Sprintf %b? Really need a better way, this looks messy.
	linkFlagBits := fmt.Sprintf("%b", lf)
	for bitIndex := 0; bitIndex < len(linkFlagBits); bitIndex++ {
		// Because we are converting to string, we need to compare the byte
		// to 0x31 which is 1 in ASCII-Hex.
		if linkFlagBits[bitIndex] == 0x31 {
			flags = append(flags, linkFlags[bitIndex])
		}
	}
	head.LinkFlags = flags
	// fmt.Println(flags)

	// CreationTime - AccessTime - WriteTime
	// var crTime, acTime, wrTime uint64
	// err = binary.Read(buf, binary.LittleEndian, &crTime)
	// if err != nil {
	// 	return head, fmt.Errorf("lnk.header: error reading CreationTime - %s", err.Error())
	// }
	// err = binary.Read(buf, binary.LittleEndian, &acTime)
	// if err != nil {
	// 	return head, fmt.Errorf("lnk.header: error reading AccessTime - %s", err.Error())
	// }
	// err = binary.Read(buf, binary.LittleEndian, &wrTime)
	// if err != nil {
	// 	return head, fmt.Errorf("lnk.header: error reading WriteTime - %s", err.Error())
	// }

	// fmt.Printf("%x\n", crTime)
	// fmt.Printf("%v\n", crTime)
	// t := time.Unix(0, (int64(crTime)/10000000)-11644473600)
	// fmt.Println(t.String())

	var crTime [16]byte
	err = binary.Read(buf, binary.LittleEndian, &crTime)
	if err != nil {
		return head, fmt.Errorf("lnk.header: error reading CreationTime - %s", err.Error())
	}

	// https://golang.org/src/syscall/types_windows.go#L344 to the rescue.
	// TODO: Convert to function.
	// TODO: Document this somewhere.
	ft := &syscall.Filetime{
		LowDateTime:  binary.LittleEndian.Uint32(crTime[:8]),
		HighDateTime: binary.LittleEndian.Uint32(crTime[8:]),
	}

	t := time.Unix(0, ft.Nanoseconds()).Format("2006-01-02 15:04:05.999999 -07:00")
	fmt.Println(t)

	return head, err
}

// // timeFromWinToUnix converts a Windows style timestamp to Unix.
// func timeFromWinToUnix(t uint64) t {

// }
