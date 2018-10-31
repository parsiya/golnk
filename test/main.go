package main

import (
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
	// lnkBytes, err := ioutil.ReadAll(fi)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Read %d bytes.\n", len(lnkBytes))

	h, err := lnk.Header(fi)
	if err != nil {
		panic(err)
	}
	_ = h

	// fmt.Println(h)

	li, err := lnk.LinkTarget(fi)
	if err != nil {
		panic(err)
	}
	_ = li
}
