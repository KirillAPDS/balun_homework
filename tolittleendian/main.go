package main

import (
	"fmt"
	"unsafe"
)

func ToLittleEndian[T ~uint16 | ~uint32 | ~uint64](num T) T {
	size := int(unsafe.Sizeof(num))
	pointer := unsafe.Pointer(&num)

	for i := 0; i < size / 2; i++ {
		first := (*int8)(unsafe.Add(pointer, i))
		last := (*int8)(unsafe.Add(pointer, size-i-1))

		*first, *last = *last, *first
	}

	return num
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