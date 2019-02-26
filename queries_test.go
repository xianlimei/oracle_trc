package main

import (
	"reflect"
	"testing"
)

func Test_detectQuery(t *testing.T) {
	type args struct {
		pl []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Detect low case select",
			args: args{
				pl: []uint8{0x0, 0xc3, 0x0, 0x0, 0x6, 0x0, 0x0, 0x0, 0x0, 0x0, 0x11, 0x69, 0xd, 0x1, 0x1, 0x1, 0x1, 0x1, 0x3,
					0x5e, 0xe, 0x2, 0x80, 0x61, 0x0, 0x1, 0x2, 0x1, 0x6b, 0x1, 0x1, 0xd, 0x1, 0x1, 0x0, 0x1, 0x64, 0x0, 0x0, 0x0,
					0x0, 0x1, 0x0, 0x1, 0x1, 0x1, 0x0, 0x0, 0x1, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0xfe, 0x40, 0x73, 0x65, 0x6c, 0x65,
					0x63, 0x74, 0x20, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x28, 0x2a, 0x29, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x75, 0x73,
					0x65, 0x72, 0x5f, 0x74, 0x61, 0x62, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x20, 0x77, 0x68, 0x65, 0x72,
					0x65, 0x20, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x3d, 0x20, 0x4e, 0x27, 0x42, 0x57,
					0x5f, 0x55, 0x53, 0x39, 0x45, 0x52, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x49,
					0x4f, 0x4e, 0x27, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
					0x20, 0x3d, 0x20, 0x4e, 0x27, 0x55, 0x41, 0x55, 0x5f, 0x46, 0x4f, 0x52, 0x47, 0x4f, 0x54, 0x5f, 0x41, 0x43, 0x54,
					0x49, 0x4f, 0x4e, 0x27, 0x0, 0x1, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0},
			},
			want: 56,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectQuery(tt.args.pl); got != tt.want {
				t.Errorf("detectQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toUpperAscii(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "All caps",
			args: args{
				b: []byte("ALL CAPS"),
			},
			want: []byte("ALL CAPS"),
		},
		{
			name: "All small",
			args: args{
				b: []byte("all small"),
			},
			want: []byte("ALL SMALL"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toUpperAscii(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toUpperAscii() = %v, want %v", got, tt.want)
			}
		})
	}
}
