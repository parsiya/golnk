package lnk

import (
	"encoding/binary"
	"syscall"
	"time"
)

// toTime converts an 8-byte Windows Filetime to time.Time. If the first
func toTime(t [8]byte) time.Time {
	// https://golang.org/src/syscall/types_windows.go#L344 to the rescue.
	ft := &syscall.Filetime{
		LowDateTime:  binary.LittleEndian.Uint32(t[:4]),
		HighDateTime: binary.LittleEndian.Uint32(t[4:]),
	}
	return time.Unix(0, ft.Nanoseconds())
}

// formatTime converts a 8-byte Windows Filetime to time.Time and then formats
// it to string.
func formatTime(t [8]byte) string {
	return toTime(t).Format("2006-01-02 15:04:05.999999 -07:00")
}
