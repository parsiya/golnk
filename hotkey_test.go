package lnk

import (
	"encoding/binary"
	"testing"
)

func TestHotKey(t *testing.T) {
	type args struct {
		hotkey uint32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"shift+0", args{byteme(0x30, 0x01)}, "SHIFT+0"},
		{"shift+Z", args{byteme(0x5A, 0x01)}, "SHIFT+Z"},
		{"invalid-low", args{byteme(0x01, 0x01)}, ""},
		{"invalid-low", args{byteme(0x00, 0x01)}, ""},
		{"invalid-high", args{byteme(0x35, 0x00)}, ""},
		{"invalid-high", args{byteme(0x35, 0x05)}, ""},
		{"alt+F12", args{byteme(0x7B, 0x04)}, "ALT+F12"},
		{"ctrl+F12", args{byteme(0x7B, 0x02)}, "CTRL+F12"},
		{"invalid-low-between", args{byteme(0x5B, 0x02)}, ""},
		{"invalid-low-between", args{byteme(0x69, 0x02)}, ""},
		{"invalid-low-over", args{byteme(0x69, 0x02)}, ""},
		{"alt+numlock", args{byteme(0x90, 0x04)}, "ALT+NUM LOCK"},
		{"shift+scrolllock", args{byteme(0x91, 0x01)}, "SHIFT+SCROLL LOCK"},
		{"invalid-low-over", args{byteme(0xFF, 0x01)}, ""},
		{"invalid-both-over", args{byteme(0xFF, 0x10)}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HotKey(tt.args.hotkey); got != tt.want {
				t.Errorf("HotKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

// convert two-byte slice to uint32.
func byteme(lb, hb uint8) uint32 {
	b := []byte{lb, hb, 0x00, 0x00}
	return binary.LittleEndian.Uint32(b)
}

func TestByteMask(t *testing.T) {
	type args struct {
		b uint32
		n int
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"byte-0", args{b: 0xDDCCBBAA, n: 0}, 0x000000AA},
		{"byte-1", args{b: 0xDDCCBBAA, n: 1}, 0x000000BB},
		{"byte-2", args{b: 0xDDCCBBAA, n: 2}, 0x000000CC},
		{"byte-3", args{b: 0xDDCCBBAA, n: 3}, 0x000000DD},
		{"invalid", args{b: 0x04030201, n: 0}, 0x00000001},
		{"byte-3", args{b: 0x04030201, n: 1}, 0x00000002},
		{"byte-3", args{b: 0x04030201, n: 2}, 0x00000003},
		{"byte-3", args{b: 0x04030201, n: 3}, 0x00000004},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteMask(tt.args.b, tt.args.n); got != tt.want {
				t.Errorf("ByteMask() = %v, want %v", got, tt.want)
			}
		})
	}
}
