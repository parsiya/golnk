package lnk

import (
	"testing"
)

func TestHotKey(t *testing.T) {
	type args struct {
		hotkey uint16
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"shift+0", args{uint16(0x0130)}, "SHIFT+0"},
		{"shift+Z", args{uint16(0x015A)}, "SHIFT+Z"},
		{"invalid-low", args{uint16(0x0101)}, "No Hotkey Set"},
		{"invalid-low", args{uint16(0x0001)}, "No Hotkey Set"},
		{"invalid-high", args{uint16(0x0035)}, "No Hotkey Set"},
		{"invalid-high", args{uint16(0x0535)}, "No Hotkey Set"},
		{"alt+F12", args{uint16(0x047B)}, "ALT+F12"},
		{"ctrl+F12", args{uint16(0x027B)}, "CTRL+F12"},
		{"invalid-low-between", args{uint16(0x025B)}, "No Hotkey Set"},
		{"invalid-low-between", args{uint16(0x0269)}, "No Hotkey Set"},
		{"invalid-low-over", args{uint16(0x0269)}, "No Hotkey Set"},
		{"alt+numlock", args{uint16(0x0490)}, "ALT+NUM LOCK"},
		{"shift+scrolllock", args{uint16(0x0191)}, "SHIFT+SCROLL LOCK"},
		{"invalid-low-over", args{uint16(0x01FF)}, "No Hotkey Set"},
		{"invalid-both-over", args{uint16(0x10FF)}, "No Hotkey Set"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HotKey(tt.args.hotkey); got != tt.want {
				t.Errorf("HotKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteMask(t *testing.T) {
	type args struct {
		b uint16
		n int
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{"byte-0", args{b: 0xBBAA, n: 0}, 0x00AA},
		{"byte-1", args{b: 0xBBAA, n: 1}, 0x00BB},
		{"invalid", args{b: 0x0201, n: 0}, 0x0001},
		{"byte-1", args{b: 0x0201, n: 1}, 0x0002},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteMask(tt.args.b, tt.args.n); got != tt.want {
				t.Errorf("ByteMask() = %v, want %v", got, tt.want)
			}
		})
	}
}
