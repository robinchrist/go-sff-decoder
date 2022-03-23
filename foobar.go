package main

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please specify decode type and file path!")
		os.Exit(1)
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()

	args := os.Args[1:]
	data, err := os.ReadFile(args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read %v bytes\n", len(data))

	if args[0] == "8636" {
		if len(data) < 256 {
			log.Fatalf("Data must have at least 256 bytes for mode 8636, but has %v", len(data))
		}
		fmt.Printf("Should decode as: %v\n", DecodeWhat(data))
		fmt.Printf("8636 decoded: %+v\n", decode8636(data))
	} else if args[0] == "8636_upper00h" {
		if len(data) < 128 {
			log.Fatalf("Data must have at least 128 bytes for mode 8636_upper00h, but has %v", len(data))
		}
		datanew := make([]byte, 128+len(data))
		for i := 0; i < 128; i++ {
			datanew[i] = 0x0
		}
		copy(datanew[128:], data)

		fmt.Printf("Should decode as: %v\n", DecodeWhat(datanew))
		fmt.Printf("8636 upper 00h decoded: %+v\n", decode8636(datanew))

	} else if args[0] == "8079" {
		if len(data) < 128 {
			log.Fatalf("Data must have at least 128 bytes for mode 8079, but has %v", len(data))
		}
		fmt.Printf("Should decode as: %v\n", DecodeWhat(data))
		fmt.Printf("8079 decoded: %+v\n", decode8079(data))
	} else {
		log.Fatalf("Invalid mode %s", args[0])
	}

	//foobar, _ := json.Marshal(decode8636(data))
	//fmt.Printf("foobar: %v\n", string(foobar))
	//fmt.Printf("%08b\n", 0b11111000)
	//fmt.Printf("%08b\n", 0x8&0x7)
}
