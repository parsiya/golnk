package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/parsiya/golnk"
)

// Sample program to parse all lnk files in the "All Users" start menu at
// C:\ProgramData\Microsoft\Windows\Start Menu\Programs.

func main() {
	startMenu := "C:/ProgramData/Microsoft/Windows/Start Menu/Programs"

	basePaths := []string{}

	err := filepath.Walk(startMenu, func(path string, info os.FileInfo, walkErr error) error {
		// Only look for lnk files.
		if filepath.Ext(info.Name()) == ".lnk" {
			f, lnkErr := lnk.File(path)
			// Print errors and move on to the next file.
			if lnkErr != nil {
				fmt.Println(lnkErr)
				return nil
			}
			var targetPath = ""
			if f.LinkInfo.LocalBasePath != "" {
				targetPath = f.LinkInfo.LocalBasePath
			}
			if f.LinkInfo.LocalBasePathUnicode != "" {
				targetPath = f.LinkInfo.LocalBasePathUnicode
			}
			if targetPath != "" {
				fmt.Println("Found", targetPath)
				basePaths = append(basePaths, targetPath)
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	// Print everything.
	fmt.Println("------------------------")
	for _, p := range basePaths {
		fmt.Println(p)
	}
}
