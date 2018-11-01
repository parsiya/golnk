package lnk

import (
	"bytes"
	"encoding/binary"
	"fmt"
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

// uint32Byte converts a uint32 to a []byte.
func uint32Byte(u uint32) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, u)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}
