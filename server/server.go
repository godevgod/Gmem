package main

import (
	"fmt"
	"log"
	"syscall"
	"time"
	"unsafe"
)

//note: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-getlargepageminimum

func main() {

	file, _ := syscall.UTF16PtrFromString("ShareMemory")

	size := 100000 //I've tried unsafe.Sizeof(TestData{}) but that didn't work.

	handle, err := syscall.CreateFileMapping(0, nil, syscall.PAGE_READWRITE, 0, uint32(size), file)

	if err != nil {

		log.Fatal(err)

	}

	defer syscall.CloseHandle(handle)

	addr, err := syscall.MapViewOfFile(handle, syscall.FILE_MAP_WRITE, 0, 0, 0)

	if err != nil {

		log.Fatal(err)

	}

	//var i byte = 0x30

	var i byte = 0x01

	for {

		data := (*Test)(unsafe.Pointer(addr))

		data.Context[0] = i

		data.str[0] = "iam"

		i++

		data.Context[1] = i

		data.str[1] = "GoDevGod"

		time.Sleep(1 * time.Second)

		//fmt.Printf("ava %v cam %v id %v\n", data.Avatar.Position, data.Camera, data.Identity)

		fmt.Printf("str: %s\n", string(data.Context[:]))
		fmt.Printf("var1 = %T\n", data.Context[:])
		println(data.str[0:0])
		fmt.Printf("var1 = %T\n", data.str[:])

		i++

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
