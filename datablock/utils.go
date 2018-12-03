package datablock

import "fmt"

// uint32Str converts a uint32 to string using fmt.Sprint.
func uint32Str(u uint32) string {
	return fmt.Sprint(u)
}

// uint32StrHex converts a uint32 to a hex encoded string using fmt.Sprintf.
func uint32StrHex(u uint32) string {
	str := fmt.Sprintf("%x", u)
	// Add a 0 to the start of odd-length string. This converts "0x1AB" to "0x01AB"
	if (len(str) % 2) != 0 {
		str = "0" + str
	}
	return "0x" + str
}

// uint32TableStr creates a string that has both decimal and hex values
// of uint32.
func uint32TableStr(u uint32) string {
	return fmt.Sprintf("%s - %s", uint32Str(u), uint32StrHex(u))
}
