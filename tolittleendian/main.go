package main

import (
	"fmt"
)

type UintAny interface {
	~uint16 | ~uint32 | ~uint64
}

func ToLittleEndian[T UintAny](x T) T {
	switch any(x).(type) {
	case uint16:
		return T(reverse16(uint16(x)))
	case uint32:
		return T(reverse32(uint32(x)))
	case uint64:
		return T(reverse64(uint64(x)))
	default:
		return x
	}
}

func reverse16(x uint16) uint16 {
	return (x>>8)|(x<<8)
}
func reverse32(x uint32) uint32 {
	return (x>>24) | ((x>>8)&0x0000FF00) | ((x<<8)&0x00FF0000) | (x<<24)
}
func reverse64(x uint64) uint64 {
	return (x>>56) |
		((x>>40)&0x000000000000FF00) |
		((x>>24)&0x0000000000FF0000) |
		((x>>8) &0x00000000FF000000) |
		((x<<8) &0x000000FF00000000) |
		((x<<24)&0x0000FF0000000000) |
		((x<<40)&0x00FF000000000000) |
		(x<<56)
}

func main() {
	var (
		a uint16 = 0x00FF
		b uint32 = 0x0000FFFF
		c uint64 = 0x00000000000000FF
		d uint16 = 0xABCD
		e uint32 = 0xAABBCDEF
		f uint64 = 0xFF010203040506AA
	)

	fmt.Printf("0x%04X\n", ToLittleEndian(a))
	fmt.Printf("0x%08X\n", ToLittleEndian(b))
	fmt.Printf("0x%016X\n", ToLittleEndian(c))
	fmt.Printf("0x%04X\n", ToLittleEndian(d))
	fmt.Printf("0x%08X\n", ToLittleEndian(e))
	fmt.Printf("0x%016X\n", ToLittleEndian(f))
}