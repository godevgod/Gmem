//READER

package main

import (
	"bytes"
	"fmt"
	"log"
	"syscall"
	"time"
	"unsafe"
)

const (
	size = 100000 // Define the size of the memory map, should match the writer's size.
)

type Test struct {
	Context [256]byte
	Str1    [256]byte
	Str2    [256]byte
}

func main() {
	file, err := syscall.UTF16PtrFromString("ShareMemory")
	if err != nil {
		log.Fatalf("Error converting string to UTF16: %v", err)
	}

	handle, err := syscall.CreateFileMapping(syscall.InvalidHandle, nil, syscall.PAGE_READONLY, 0, uint32(size), file)
	if err != nil {
		log.Fatalf("Error creating file mapping: %v", err)
	}
	defer syscall.CloseHandle(handle)

	addr, err := syscall.MapViewOfFile(handle, syscall.FILE_MAP_READ, 0, 0, 0)
	if err != nil {
		log.Fatalf("Error mapping view of file: %v", err)
	}
	defer syscall.UnmapViewOfFile(addr)

	data := (*Test)(unsafe.Pointer(addr))

	for {
		// Read and convert byte arrays back to strings
		str1 := string(bytes.Trim(data.Str1[:], "\x00"))
		str2 := string(bytes.Trim(data.Str2[:], "\x00"))

		// Print the strings
		fmt.Printf("Str1: %s, Str2: %s\n", str1, str2)

		// Sleep for a second before the next read
		time.Sleep(1 * time.Second)
	}
}
