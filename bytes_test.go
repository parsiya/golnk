package lnk

import (
	"testing"
)

var (
	b0 = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07}
	b1 = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A}
	b2 = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D}
	b3 = []byte{0xFF, 0xEE, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D}
	b4 = []byte{0xCC, 0xDD, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D}
	b5 = []byte{0xBB, 0xAA, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D}
)

func Test_uint16Little(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{"b0", args{b0}, 0x0100},
		{"b1", args{b1}, 0x0100},
		{"b2", args{b2}, 0x0100},
		{"b3", args{b3}, 0xEEFF},
		{"b4", args{b4}, 0xDDCC},
		{"b5", args{b5}, 0xAABB},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uint16Little(tt.args.b); got != tt.want {
				t.Errorf("uint16Little() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uint32Little(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"b0", args{b0}, 0x03020100},
		{"b1", args{b1}, 0x03020100},
		{"b2", args{b2}, 0x03020100},
		{"b3", args{b3}, 0x0302EEFF},
		{"b4", args{b4}, 0x0302DDCC},
		{"b5", args{b5}, 0x0302AABB},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uint32Little(tt.args.b); got != tt.want {
				t.Errorf("uint32Little() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uint64Little(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"b0", args{b0}, 0x0706050403020100},
		{"b1", args{b1}, 0x0706050403020100},
		{"b2", args{b2}, 0x0706050403020100},
		{"b3", args{b3}, 0x070605040302EEFF},
		{"b4", args{b4}, 0x070605040302DDCC},
		{"b5", args{b5}, 0x070605040302AABB},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uint64Little(tt.args.b); got != tt.want {
				t.Errorf("uint64Little() = %v, want %v", got, tt.want)
			}
		})
	}
}
