package lnk

import (
	"testing"
	"time"
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

func Test_toTime(t *testing.T) {
	type args struct {
		ft [8]byte // filetime as read from disk
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "zero-time",
			args: args{ft: [8]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
			want: time.Time{},
		},
		{
			name: "unix-epoch",
			args: args{ft: [8]byte{0x00, 0x80, 0x3E, 0xD5, 0xDE, 0xB1, 0x9D, 0x01}},
			want: time.Unix(0, 0).UTC(),
		},
		{
			name: "year-2000",
			args: args{ft: [8]byte{0x00, 0xB0, 0x07, 0x70, 0x1D, 0x54, 0xBF, 0x01}},
			want: time.Date(2000, 1, 1, 6, 0, 0, 0, time.UTC),
		},
		{
			name: "year-2023",
			args: args{ft: [8]byte{0x00, 0x70, 0x5D, 0x48, 0xA6, 0x1D, 0xD9, 0x01}},
			want: time.Date(2023, 1, 1, 6, 0, 0, 0, time.UTC),
		},
		{
			name: "with-time-components",
			args: args{ft: [8]byte{0x80, 0xF8, 0x46, 0x29, 0x0F, 0x1E, 0xD9, 0x01}},
			want: time.Date(2023, 1, 1, 18, 30, 45, 0, time.UTC),
		},
		{
			name: "with-microseconds",
			args: args{ft: [8]byte{0x40, 0xB2, 0x33, 0xA6, 0xF1, 0x4C, 0xA3, 0x01}},
			want: time.Date(1975, 1, 1, 6, 0, 0, 100000000, time.UTC),
		},
		{
			name: "leap-year-feb29",
			args: args{ft: [8]byte{0x00, 0xF0, 0x66, 0x36, 0x7A, 0x82, 0xBF, 0x01}},
			want: time.Date(2000, 2, 29, 6, 0, 0, 0, time.UTC),
		},
		{
			name: "end-of-day",
			args: args{ft: [8]byte{0xC0, 0x2D, 0x62, 0x9A, 0xE6, 0x54, 0xBF, 0x01}},
			want: time.Date(2000, 1, 2, 5, 59, 59, 900000000, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toTime(tt.args.ft)
			if !got.Equal(tt.want) {
				t.Errorf("%s: toTime() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
