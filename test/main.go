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

	h, err := lnk.Header(lnkBytes)
	if err != nil {
		panic(err)
	}

	fmt.Println(h)

}
