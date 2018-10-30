package main

import (
	"fmt"
	"io/ioutil"
	"os"

	lnk "github.com/parsiya/golnk"
)

func main() {

	fi, err := os.Open("test.lnk")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	// lnk files are small-ish, no reason not to read everything at once.
	lnkBytes, err := ioutil.ReadAll(fi)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read %d bytes.\n", len(lnkBytes))

	// buf := bytes.NewReader(lnkBytes)
	// // First four bytes == header. Must be 0x4c in little-endian.
	// var header uint32
	// err = binary.Read(buf, binary.LittleEndian, &header)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%x\n", header)

	// if header != 0x4c {
	// 	fmt.Printf("Bad header - got %x - wanted %x", header, 0x4c)
	// }

	_, err = lnk.Header(lnkBytes)
	if err != nil {
		panic(err)
	}

}
