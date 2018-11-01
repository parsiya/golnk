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
		{"invalid-low", args{uint16(0x0101)}, "No Key Assigned"},
		{"invalid-low", args{uint16(0x0001)}, "No Key Assigned"},
		{"invalid-high", args{uint16(0x0035)}, "No Key Assigned"},
		{"invalid-high", args{uint16(0x0535)}, "No Key Assigned"},
		{"alt+F12", args{uint16(0x047B)}, "ALT+F12"},
		{"ctrl+F12", args{uint16(0x027B)}, "CTRL+F12"},
		{"invalid-low-between", args{uint16(0x025B)}, "No Key Assigned"},
		{"invalid-low-between", args{uint16(0x0269)}, "No Key Assigned"},
		{"invalid-low-over", args{uint16(0x0269)}, "No Key Assigned"},
		{"alt+numlock", args{uint16(0x0490)}, "ALT+NUM LOCK"},
		{"shift+scrolllock", args{uint16(0x0191)}, "SHIFT+SCROLL LOCK"},
		{"invalid-low-over", args{uint16(0x01FF)}, "No Key Assigned"},
		{"invalid-both-over", args{uint16(0x10FF)}, "No Key Assigned"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HotKey(tt.args.hotkey); got != tt.want {
				t.Errorf("HotKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
