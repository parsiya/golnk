package lnk

import (
	"encoding/binary"
	"fmt"
)

// ByteMask returns one of the four bytes from a uint32.
func ByteMask(b uint32, n int) uint32 {
	// Maybe we should not panic, hmm.
	if n < 0 || n > 3 {
		panic(fmt.Sprintf("invalid byte mask, got %d", n))
	}
	mask := uint32(0x000000FF) << uint32(n*8)
	return (b & mask) >> uint32(n*8)
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

// uint32Str converts a uint32 to string using fmt.Sprint.
func uint32Str(u uint32) string {
	return fmt.Sprint(u)
}

// int32Str converts a int32 to string using fmt.Sprint.
func int32Str(u int32) string {
	return fmt.Sprint(u)
}
