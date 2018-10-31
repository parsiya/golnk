package lnk

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/olekukonko/tablewriter"
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
	head.Magic = magic

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
	// We will lose preceding zeros by using uint32 but we do not care about them.
	var lf uint32
	err = binary.Read(buf, binary.LittleEndian, &lf)
	if err != nil {
		return head, fmt.Errorf("lnk.header: reading LinkFlags - %s", err.Error())
	}
	flags = matchFlag(lf, linkFlags)
	head.LinkFlags = flags

	// Parse FileAttributes.
	var attribs uint32
	// Same as before, read BigEndian.
	err = binary.Read(buf, binary.LittleEndian, &attribs)
	if err != nil {
		return head, fmt.Errorf("lnk.header: reading FileAttributes - %s", err.Error())
	}
	head.FileAttributes = matchFlag(attribs, fileAttributesFlags)

	// Convert timestamps from Windows Filetime to time.Time.
	var crTime, wrTime, acTime [8]byte
	err = binary.Read(buf, binary.LittleEndian, &crTime)
	if err != nil {
		return head, fmt.Errorf("lnk.header: reading CreationTime - %s", err.Error())
	}
	head.CreationTime = toTime(crTime)

	err = binary.Read(buf, binary.LittleEndian, &wrTime)
	if err != nil {
		return head, fmt.Errorf("lnk.header: reading WriteTime - %s", err.Error())
	}
	head.WriteTime = toTime(wrTime)

	err = binary.Read(buf, binary.LittleEndian, &acTime)
	if err != nil {
		return head, fmt.Errorf("lnk.header: reading AccessTime - %s", err.Error())
	}
	head.AccessTime = toTime(acTime)

	// Target file size.
	var size uint32
	err = binary.Read(buf, binary.LittleEndian, &size)
	if err != nil {
		return head, fmt.Errorf("lnk.header: reading target file size - %s", err.Error())
	}
	head.TargetFileSize = size

	// Icon index is a signed 32-bit integer.
	var iconIndex int32
	err = binary.Read(buf, binary.LittleEndian, &iconIndex)
	if err != nil {
		return head, fmt.Errorf("lnk.header: reading icon index - %s", err.Error())
	}
	head.IconIndex = iconIndex

	// ShowCommand
	var sw uint32
	err = binary.Read(buf, binary.LittleEndian, &sw)
	if err != nil {
		return head, fmt.Errorf("lnk.header: reading showcommand - %s", err.Error())
	}
	head.ShowCommand = showCommand(sw)

	// Hotkey.
	var hk uint32
	err = binary.Read(buf, binary.LittleEndian, &hk)
	if err != nil {
		return head, fmt.Errorf("lnk.header: reading hotkey - %s", err.Error())
	}
	head.HotKey = HotKey(hk)

	// The rest should be 10-bytes of zeroes.
	binary.Read(buf, binary.LittleEndian, &head.Reserved1)
	binary.Read(buf, binary.LittleEndian, &head.Reserved2)
	binary.Read(buf, binary.LittleEndian, &head.Reserved3)

	return head, err
}

// String prints the ShellLinkHeader in a nice looking table.
func (h ShellLinkHeader) String() string {
	var sb, flags, attribs strings.Builder

	// Append all flags.
	for _, fl := range h.LinkFlags {
		flags.WriteString(fl)
		flags.WriteString("\n")
	}

	// Append all file attributes.
	for _, at := range h.FileAttributes {
		attribs.WriteString(at)
		attribs.WriteString("\n")
	}

	table := tablewriter.NewWriter(&sb)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetRowLine(true)

	table.SetHeader([]string{"ShellLinkHeader Field", "Value"})

	table.Append([]string{"Magic", uint32Str(h.Magic)})
	table.Append([]string{"LinkCLSID", hex.EncodeToString(h.LinkCLSID[:])})
	table.Append([]string{"LinkFlags", flags.String()})
	table.Append([]string{"FileAttributes", attribs.String()})
	table.Append([]string{"CreationTime", h.CreationTime.String()})
	table.Append([]string{"AccessTime", h.AccessTime.String()})
	table.Append([]string{"WriteTime", h.WriteTime.String()})
	table.Append([]string{"TargetFileSize", uint32Str(h.TargetFileSize)})
	table.Append([]string{"IconIndex", int32Str(h.IconIndex)})
	table.Append([]string{"HotKey", h.HotKey})
	table.Render()

	return sb.String()
}
