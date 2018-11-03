package main

import (
	"os"

	lnk "github.com/parsiya/golnk"
)

func main() {

	fi, err := os.Open("remote.directory.xp.test")
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

	// fmt.Println(lnk.StructToJSON(h, true))

	li, err := lnk.LinkTarget(fi)
	if err != nil {
		panic(err)
	}
	_ = li
	// fmt.Println(lnk.StructToJSON(li, true))

	_, err = lnk.LinkInfo(fi)
	if err != nil {
		panic(err)
	}

}
