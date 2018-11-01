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

func Test_byteMaskuint16(t *testing.T) {
	type args struct {
		b uint16
		n int
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{"highbyte", args{uint16(0xAABB), 1}, 0xAA},
		{"lowbyte", args{uint16(0xAABB), 0}, 0xBB},
		{"highbyte-zero", args{uint16(0x00BB), 1}, 0x00},
		{"lowbyte-zero", args{uint16(0xAA00), 0}, 0x00},
		{"byte-0", args{b: 0xBBAA, n: 0}, 0x00AA},
		{"byte-1", args{b: 0xBBAA, n: 1}, 0x00BB},
		{"invalid", args{b: 0x0201, n: 0}, 0x0001},
		{"byte-1", args{b: 0x0201, n: 1}, 0x0002},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := byteMaskuint16(tt.args.b, tt.args.n); got != tt.want {
				t.Errorf("byteMaskuint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bitMaskuint32(t *testing.T) {
	// 255 is 1111 1111
	u255 := uint32(255)
	// 0 is 0000 0000
	u0 := uint32(0)
	type args struct {
		b uint32
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"bit0-255", args{u255, 0}, true},
		{"bit1-255", args{u255, 0}, true},
		{"bit2-255", args{u255, 0}, true},
		{"bit3-255", args{u255, 0}, true},
		{"bit4-255", args{u255, 0}, true},
		{"bit5-255", args{u255, 0}, true},
		{"bit6-255", args{u255, 0}, true},
		{"bit7-255", args{u255, 0}, true},
		{"bit0-0", args{u0, 0}, false},
		{"bit1-0", args{u0, 0}, false},
		{"bit2-0", args{u0, 0}, false},
		{"bit3-0", args{u0, 0}, false},
		{"bit4-0", args{u0, 0}, false},
		{"bit5-0", args{u0, 0}, false},
		{"bit6-0", args{u0, 0}, false},
		{"bit7-0", args{u0, 0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitMaskuint32(tt.args.b, tt.args.n); got != tt.want {
				t.Errorf("bitMaskuint32() = %v, want %v", got, tt.want)
			}
		})
	}
}
