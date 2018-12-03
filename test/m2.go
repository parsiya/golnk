package main

import (
	"fmt"

	"github.com/parsiya/golnk"
)

func main() {

	Lnk, err := lnk.File("test.lnk")
	if err != nil {
		panic(err)
	}

	// Print header.
	fmt.Println(Lnk.Header)

	// Print LocalBasePath.
	fmt.Println("BasePath", Lnk.LinkInfo.LocalBasePath)

	fmt.Println(Lnk.LinkInfo)

	fmt.Println(Lnk.StringData)

	fmt.Println(Lnk.DataBlocks)
}
