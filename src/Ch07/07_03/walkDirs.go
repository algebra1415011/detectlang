package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

var (
	cnt int
	ext string
)

func main() {

	start := time.Now()
	root, _ := filepath.Abs("C:/Users/admin/Desktop/Go/Rx")
	fmt.Println("Processing path", root)
	ext = ".java"
	err := filepath.Walk(root, processPath)
	elapsed := time.Since(start)
	fmt.Println("Source code took %s and count of file %v", elapsed, cnt)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func processPath(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if path != "." {
		if info.IsDir() {
			fmt.Println("Directory:", path)
		} else {
			fmt.Println("File:", path)
			if checkExt(info.Name()) {
				// fmt.Println("File:", info.Name())
				cnt++
			}
		}
	}

	return nil
}

func checkExt(filename string) bool {
	r, err := regexp.MatchString(ext, filename)
	if err == nil && r {
		return true
	}
	return false

}
