package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v

func TestConversion(t *testing.T) {
	type test struct {
		name	string
		have  	any
		want 	any
	}

	cases := []test{

		{"uint16 #1", uint16(0x0000), uint16(0x0000)},
		{"uint16 #2", uint16(0x00FF), uint16(0xFF00)},
		{"uint16 #3", uint16(0x1234), uint16(0x3412)},
		{"uint16 #4", uint16(0xABCD), uint16(0xCDAB)},
		{"uint16 #5", uint16(0xDADC), uint16(0xDCDA)},
		{"uint16 #6", uint16(0xFFFF), uint16(0xFFFF)},

		{"uint32 #1", uint32(0x00000000), uint32(0x00000000)},
		{"uint32 #2", uint32(0xFFFFFFFF), uint32(0xFFFFFFFF)},
		{"uint32 #3", uint32(0x00FF00FF), uint32(0xFF00FF00)},
		{"uint32 #4", uint32(0x0000FFFF), uint32(0xFFFF0000)},
		{"uint32 #5", uint32(0x01020304), uint32(0x04030201)},
		{"uint32 #6", uint32(0xAABBCDEF), uint32(0xEFCDBBAA)},

		{"uint64 #1", uint64(0x0000000000000000), uint64(0x0000000000000000)},
		{"uint64 #2", uint64(0xFFFFFFFFFFFFFFFF), uint64(0xFFFFFFFFFFFFFFFF)},
		{"uint64 #3", uint64(0x00000000000000FF), uint64(0xFF00000000000000)},
		{"uint64 #4", uint64(0x0102030405060708), uint64(0x0807060504030201)},
		{"uint64 #5", uint64(0xFF010203040506AA), uint64(0xAA060504030201FF)},
		{"uint64 #6", uint64(0xAABBCCDDEEFF1020), uint64(0x2010FFEEDDCCBBAA)},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			switch v := c.have.(type) {
			case uint16:
				got := ToLittleEndian(v)
				assert.Equal(t, c.want.(uint16), got, "uint16 mismatch")
			case uint32:
				got := ToLittleEndian(v)
				assert.Equal(t, c.want.(uint32), got, "uint32 mismatch")
			case uint64:
				got := ToLittleEndian(v)
				assert.Equal(t, c.want.(uint64), got, "uint64 mismatch")
			default:
				t.Fatalf("unsupported type: %T", c.have)
			}
		})
	}
}

