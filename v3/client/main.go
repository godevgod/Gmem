//READER

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"syscall"
	"time"
	"unsafe"
)

const (
	size     = 100000 // Define the size of the memory map.
	hashSize = 8      // Size for the hash value.
)

type Test struct {
	Context [256]byte
	Str1    [256]byte
	Str2    [256]byte
	Hash    [hashSize]byte // Added field for the hash.
}

func main() {
	file, err := syscall.UTF16PtrFromString("ShareMemory")
	if err != nil {
		log.Fatalf("Error converting string to UTF16: %v", err)
	}

	//handle, _, err := syscall.CreateFileMapping(syscall.InvalidHandle, nil, syscall.PAGE_READWRITE, 0, uint32(size), file)
	handle, err := syscall.CreateFileMapping(syscall.InvalidHandle, nil, syscall.PAGE_READWRITE, 0, uint32(size), file)
	if err != nil && err != syscall.ERROR_ALREADY_EXISTS {
		log.Fatalf("Error creating/opening file mapping: %v", err)
	}

	addr, err := syscall.MapViewOfFile(handle, syscall.FILE_MAP_READ, 0, 0, uintptr(size))
	if err != nil {
		log.Fatalf("Error mapping view of file: %v", err)
	}
	defer syscall.UnmapViewOfFile(addr)

	data := (*Test)(unsafe.Pointer(addr))

	for {
		str1 := string(bytes.Trim(data.Str1[:], "\x00"))
		str2 := string(bytes.Trim(data.Str2[:], "\x00"))
		hashValue := binary.BigEndian.Uint64(data.Hash[:])

		fmt.Printf("Str1: %s, Str2: %s, Hash: %x\n", str1, str2, hashValue)

		time.Sleep(1 * time.Second)
	}
}
