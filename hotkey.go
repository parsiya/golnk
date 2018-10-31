package lnk

import (
	"strconv"
	"strings"
)

/*
	HotKeyFlags contains the hotkey.
	Although it's 4 bytes, only the first 2 bytes are used.

	First byte is LowByte.
	If between 0x30 and 0x5A inclusive, it's ASCII-hex of the key.
	If between 0x70 and 0x87 it's F(num-0x70+1) (e.g. 0x70 == F1 and 0x87 == F24).
	0x90 == NUM LOCK and 0x91 SCROLL LOCK.

	Second byte is HighByte.
	0x01: SHIFT
	0X02: CTRL
	0X03: ALT
*/
// HotKey returns the string representation of the hotkey uint32.
func HotKey(hotkey uint16) string {
	var sb strings.Builder
	lb := ByteMask(hotkey, 0) // first byte
	hb := ByteMask(hotkey, 1) // second byte

	switch hb {
	case 0x01:
		sb.WriteString("SHIFT")
	case 0x02:
		sb.WriteString("CTRL")
	case 0x04:
		sb.WriteString("ALT")
	default:
		return "No Hotkey Set"
	}

	sb.WriteString("+")

	switch {
	case 0x30 <= lb && lb <= 0x5A:
		sb.WriteString(string(lb))
	case 0x70 <= lb && lb <= 0x87:
		sb.WriteString("F" + strconv.Itoa(int(lb-0x70+1)))
	case lb == 0x90:
		sb.WriteString("NUM LOCK")
	case lb == 0x91:
		sb.WriteString("SCROLL LOCK")
	default:
		return "No Hotkey Set"
	}
	return sb.String()
}
