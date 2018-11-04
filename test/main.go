package main

import (
	"fmt"
	"os"

	lnk "github.com/parsiya/golnk"
)

func main() {

	fi, err := os.Open("remote.directory.xp.test")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	// // lnk files are small-ish, no reason not to read everything at once.
	// // lnkBytes, err := ioutil.ReadAll(fi)
	// // if err != nil {
	// // 	panic(err)
	// // }
	// // fmt.Printf("Read %d bytes.\n", len(lnkBytes))

	// h, err := lnk.Header(fi)
	// if err != nil {
	// 	panic(err)
	// }
	// _ = h

	// // fmt.Println(lnk.StructToJSON(h, true))
	// fmt.Println(h)

	// fmt.Println(h.LinkFlags)

	// lt, err := lnk.LinkTarget(fi)
	// if err != nil {
	// 	panic(err)
	// }
	// _ = lt
	// fmt.Println(lnk.StructToJSON(lt, true))

	// li, err := lnk.LinkInfo(fi)
	// if err != nil {
	// 	panic(err)
	// }
	// _ = li
	// fmt.Println(lnk.StructToJSON(li, true))

	// st, err := lnk.StringData(fi, h.LinkFlags)
	// if err != nil {
	// 	panic(err)
	// }
	// _ = st
	// fmt.Println(lnk.StructToJSON(st, true))

	// edb, err := lnk.DataBlock(fi)
	// if err != nil {
	// 	panic(err)
	// }
	// _ = edb
	// // fmt.Println(lnk.StructToJSON(edb, true))

	ln, err := lnk.Read(fi)
	if err != nil {
		panic(err)
	}
	fmt.Println(ln.Header)

}
