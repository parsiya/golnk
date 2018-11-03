package lnk

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"unicode/utf8"
)

// byteMaskuint16 returns one of the two bytes from a uint16.
func byteMaskuint16(b uint16, n int) uint16 {
	// Maybe we should not panic, hmm.
	if n < 0 || n > 2 {
		panic(fmt.Sprintf("invalid byte mask, got %d", n))
	}
	mask := uint16(0x000000FF) << uint16(n*8)
	return (b & mask) >> uint16(n*8)
}

// bitMaskuint32 returns one of the 32-bits from a uint32.
// Returns true for 1 and false for 0.
func bitMaskuint32(b uint32, n int) bool {
	if n < 0 || n > 31 {
		panic(fmt.Sprintf("invalid bit number, got %d", n))
	}
	return ((b >> uint(n)) & 1) == 1
}

// ReadBytes reads n bytes from the slice starting from offset and
// returns a []byte and the number of bytes read. If offset is out of bounds
// it returns an empty []byte and 0 bytes read.
// TODO: Write tests for this.
func ReadBytes(b []byte, offset, num int) (out []byte, n int) {
	if offset >= len(b) {
		return out, 0
	}
	if offset+num >= len(b) {
		return b[offset:], len(b[offset:])
	}
	return b[offset : offset+num], num
}

// readSection reads a size from the start of the io.Reader. The size length is
// decided by the parameter sSize.
// sSize == 2 - read uint16
// sSize == 4 - read uint32
// sSize == 8 - read uint64 - Not needed for now.
// Then read (size-sSize) bytes, populate the start with the original bytes and
// add the rest. Finally return the []byte and a new io.Reader to it.
// The size bytes are added to the start of the []byte to keep the section []byte
// intact for later offset use.
func readSection(r io.Reader, sSize int) (data []byte, nr io.Reader, size int, err error) {
	// We are not going to lose data by copying a smaller var into a larger one.
	var sectionSize uint64
	switch sSize {
	case 2:
		// Read uint16.
		var size16 uint16
		err = binary.Read(r, binary.LittleEndian, &size16)
		if err != nil {
			return data, nr, size, fmt.Errorf("golnk.readSection: read size %d bytes - %s", sSize, err.Error())
		}
		sectionSize = uint64(size16)
		// Add bytes to the start of data []byte.
		data = uint16Byte(size16)
	case 4:
		// Read uint32.
		var size32 uint32
		err = binary.Read(r, binary.LittleEndian, &size32)
		if err != nil {
			return data, nr, size, fmt.Errorf("golnk.readSection: read size %d bytes - %s", sSize, err.Error())
		}
		sectionSize = uint64(size32)
		// Add bytes to the start of data []byte.
		data = uint32Byte(size32)
	case 8:
		// Read uint64 or sectionSize.
		err = binary.Read(r, binary.LittleEndian, &sectionSize)
		if err != nil {
			return data, nr, size, fmt.Errorf("golnk.readSection: read size %d bytes - %s", sSize, err.Error())
		}
		// Add bytes to the start of data []byte.
		data = uint64Byte(sectionSize)
	default:
		return data, nr, size, fmt.Errorf("golnk.readSection: invalid sSize - got %v", sSize)
	}

	// Create a []byte of sectionSize-4 and read that many bytes from io.Reader.
	tempData := make([]byte, sectionSize-uint64(sSize))
	err = binary.Read(r, binary.LittleEndian, &tempData)
	if err != nil {
		return data, nr, size, fmt.Errorf("golnk.readSection: read section %d bytes - %s", sectionSize-uint64(sSize), err.Error())
	}

	// If this is successful, append it to data []byte.
	data = append(data, tempData...)

	// Create a reader from the unread bytes.
	nr = bytes.NewReader(tempData)

	return data, nr, int(sectionSize), nil
}

// readString returns a string of all bytes from the []byte until the first 0x00.
// TODO: Tests for this?
func readString(data []byte) string {
	// Find the index of first 0x00.
	i := bytes.IndexByte(data, byte(0x00))
	if i == -1 {
		// If 0x00 is not found, return all the slice.
		i = len(data) - 1
	}
	return string(data[:i])
}

// readUnicodeString returns a string of all bytes from the []byte until the
// first 0x00 00.
// TODO: Write tests.
func readUnicodeString(data []byte) string {
	// Create the unicode null-terminator.
	var unicodeNull rune
	n := utf8.EncodeRune([]byte{0x00, 0x00}, unicodeNull)
	if n != 2 {
		// This failed, return an empty string.
		return ""
	}
	// Find the index of first 0x0000.
	i := bytes.IndexRune(data, unicodeNull)
	if i == -1 {
		// If 0x0000 is not found, return all the slice.
		i = len(data) - 1
	}
	return string(data[:i])
}

// uint16Little reads a uint16 from []byte and returns the result in Little-Endian.
func uint16Little(b []byte) uint16 {
	if len(b) < 2 {
		panic(fmt.Sprintf("input smaller than two bytes - got %d", len(b)))
	}
	return binary.LittleEndian.Uint16(b)
}

// uint32Little reads a uint32 from []byte and returns the result in Little-Endian.
func uint32Little(b []byte) uint32 {
	if len(b) < 4 {
		panic(fmt.Sprintf("input smaller than four bytes - got %d", len(b)))
	}
	return binary.LittleEndian.Uint32(b)
}

// uint64Little reads a uint64 from []byte and returns the result in Little-Endian.
func uint64Little(b []byte) uint64 {
	if len(b) < 8 {
		panic(fmt.Sprintf("input smaller than eight bytes - got %d", len(b)))
	}
	return binary.LittleEndian.Uint64(b)
}

// uint16Str converts a uint16 to string using fmt.Sprint.
func uint16Str(u uint16) string {
	return fmt.Sprint(u)
}

// int16Str converts an int16 to string using fmt.Sprint.
func int16Str(u int16) string {
	return fmt.Sprint(u)
}

// uint32Str converts a uint32 to string using fmt.Sprint.
func uint32Str(u uint32) string {
	return fmt.Sprint(u)
}

// int32Str converts an int32 to string using fmt.Sprint.
func int32Str(u int32) string {
	return fmt.Sprint(u)
}

// uint16Byte converts a uint16 to a []byte.
func uint16Byte(u uint16) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, u)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

// uint32Byte converts a uint32 to a []byte.
func uint32Byte(u uint32) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, u)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

// uint64Byte converts a uint64 to a []byte.
func uint64Byte(u uint64) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, u)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}
