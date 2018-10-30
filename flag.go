package lnk

import (
	"fmt"
)

// Flag utilities.

// matchFlag does the following:
// Given a uint32 flag read in littleEndian from disk and a []string,
// match the flag bits and return a []string of matched flags.
// This happens because bits are matched to the flags from 0 onwards but the
// bit string is the other way around.
func matchFlag(flag uint32, flagText []string) []string {
	// Convert to bits and then reverse.
	flagBits := Reverse(fmt.Sprintf("%b", flag))
	var fl []string
	// If we have more bits than flags (something has gone wrong or the file is corrupted),
	// then reduce the flagbits.

	if len(flagBits) > len(flagText) {
		flagBits = flagBits[:len(flagText)]
	}
	for bitIndex := 0; bitIndex < len(flagBits); bitIndex++ {
		if flagBits[bitIndex] == 0x31 {
			fl = append(fl, flagText[bitIndex])
		}
	}
	return fl
}

// Reverse returns its argument string reversed rune-wise left to right.
// Taken from https://github.com/golang/example/blob/master/stringutil/reverse.go.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
