//WRITER

package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"syscall"
	"time"
	"unsafe"
)

const (
	size           = 100000 // Define the size of the memory map.
	maxMemoryUsage = size   // maximum memory you want to use
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

	handle, err := syscall.CreateFileMapping(syscall.InvalidHandle, nil, syscall.PAGE_READWRITE, 0, uint32(size), file)
	if err != nil {
		log.Fatalf("Error creating file mapping: %v", err)
	}
	defer syscall.CloseHandle(handle)

	addr, err := syscall.MapViewOfFile(handle, syscall.FILE_MAP_WRITE, 0, 0, 0)
	if err != nil {
		log.Fatalf("Error mapping view of file: %v", err)
	}
	defer syscall.UnmapViewOfFile(addr)

	data := (*Test)(unsafe.Pointer(addr))
	rand.Seed(time.Now().UnixNano()) // Initialize the random number generator
	for {
		// Generate random data for Context
		data.Context[0] = byte(rand.Intn(256)) // Random byte between 0-255
		data.Context[1] = byte(rand.Intn(256)) // Random byte between 0-255

		// Generate random strings for Str1 and Str2
		copy(data.Str1[:], []byte(fmt.Sprintf("iam%03d", rand.Intn(1000))))      // Random string like "iam123"
		copy(data.Str2[:], []byte(fmt.Sprintf("GoDevGod%03d", rand.Intn(1000)))) // Random string like "GoDevGod456"

		// Print the updated values
		fmt.Printf("Context: [%d, %d]\n", data.Context[0], data.Context[1])
		fmt.Printf("Context: [0x%X, 0x%X]\n", data.Context[0], data.Context[1])

		str1 := string(bytes.Trim(data.Str1[:], "\x00"))
		str2 := string(bytes.Trim(data.Str2[:], "\x00"))
		fmt.Printf("Str1: %s, Str2: %s\n", str1, str2)

		// Sleep for a second before the next update
		time.Sleep(1 * time.Second)
	}
}
