package lnk

import (
	"encoding/binary"
	"fmt"
	"syscall"
	"time"
)

// toTime converts an 8-byte Windows Filetime to time.Time. If the first
func toTime(t [8]byte) time.Time {
	fmt.Printf("%x\n", t)

	lb := binary.BigEndian.Uint32(t[:4])
	hb := binary.BigEndian.Uint32(t[4:])

	fmt.Printf("lb: %x\n", lb)
	fmt.Printf("hb: %x\n", hb)

	// https://golang.org/src/syscall/types_windows.go#L344 to the rescue.
	ft := &syscall.Filetime{
		LowDateTime:  binary.BigEndian.Uint32(t[:4]),
		HighDateTime: binary.BigEndian.Uint32(t[4:]),
	}
	return time.Unix(0, ft.Nanoseconds())
}

// formatTime converts a 8-byte Windows Filetime to time.Time and then formats
// it to string.
func formatTime(t [8]byte) string {
	return toTime(t).Format("2006-01-02 15:04:05.999999 -07:00")
}
