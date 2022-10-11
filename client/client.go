package main

import (
	"fmt"
	"log"
	"syscall"
	"time"
	"unsafe"
)

func main() {

	file, _ := syscall.UTF16PtrFromString("ShareMemory")

	size := 100000 //I've tried unsafe.Sizeof(TestData{}) but that didn't work.

	handle, err := syscall.CreateFileMapping(0, nil, syscall.PAGE_READONLY, 0, uint32(size), file)

	if err != nil {

		log.Fatal(err)

	}

	defer syscall.CloseHandle(handle)

	fmt.Println(syscall.GetLastError())

	addr, err := syscall.MapViewOfFile(handle, syscall.FILE_MAP_READ, 0, 0, 0)

	if err != nil {

		log.Fatal(err)

	}

	defer syscall.UnmapViewOfFile(addr)

	for {

		data := (*Test)(unsafe.Pointer(addr))

		time.Sleep(1 * time.Second)

		//fmt.Printf("ava %v cam %v id %v\n", data.Avatar.Position, data.Camera, data.Identity)

		fmt.Printf("str: %s\n", string(data.Context[:]))
		fmt.Printf("var1 = %T\n", data.Context[:])
		println(data.str[0:0])
		fmt.Printf("var1 = %T\n", data.str[:])

	}

}

// type TestVector [3]float32
// type WChar uint16

// type TestIdentity [256]uint16

// func (id TestIdentity) String() string {
// 	buf := make([]byte, 4)
// 	var ret bytes.Buffer
// 	runes := utf16.Decode(id[:])
// 	count := 0
// 	for _, rune := range runes {
// 		utf8.EncodeRune(buf, rune)
// 		ret.WriteString(string(rune))
// 		count++
// 	}
// 	return ret.String()
// }

// type TestPosition struct {
// 	Position TestVector
// 	Front    TestVector
// 	Top      TestVector
// }

type Test struct {
	//Version       uint32
	//Tick          uint32
	//Avatar        TestPosition
	//Name          [256]WChar
	//Camera        TestPosition
	//Identity      TestIdentity
	//ContextLength uint32
	Context [256]byte
	//Description   [2048]WChar
	str [2]string
}
