//SERVER

package main

import (
	"fmt"
	"hash/fnv"
	"log"
	"math/rand"
	"syscall"
	"time"
	"unsafe"
)

const (
	size           = 100000 // Define the size of the memory map.
	maxMemoryUsage = size   // Maximum memory you want to use.
	hashSize       = 8      // Size for the hash value.
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
	rand.Seed(time.Now().UnixNano())
	hasher := fnv.New64()

	for {
		data.Context[0] = byte(rand.Intn(256))
		data.Context[1] = byte(rand.Intn(256))

		str1Value := fmt.Sprintf("iam%03d", rand.Intn(1000))
		str2Value := fmt.Sprintf("GoDevGod%03d", rand.Intn(1000))
		copy(data.Str1[:], []byte(str1Value))
		copy(data.Str2[:], []byte(str2Value))

		hasher.Reset()
		hasher.Write([]byte(str1Value))
		hasher.Write([]byte(str2Value))
		hasher.Write([]byte(fmt.Sprintf("%d", time.Now().UnixNano())))
		hashValue := hasher.Sum(nil)

		copy(data.Hash[:], hashValue)

		fmt.Printf("Context: [%d, %d]\n", data.Context[0], data.Context[1])
		fmt.Printf("Context: [0x%X, 0x%X]\n", data.Context[0], data.Context[1])
		fmt.Printf("Str1: %s, Str2: %s, Hash: %x\n", str1Value, str2Value, hashValue)

		time.Sleep(1 * time.Second)
	}
}
