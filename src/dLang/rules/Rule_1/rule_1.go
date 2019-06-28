package rule_1

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	
)

var (
	Cnt int
	ext string
)

func FindCount(extension string, projectPath string) int {

	// start := time.Now()
	// root, _ := filepath.Abs("C:/Users/admin/Desktop/Go/Rx")
	root, _ := filepath.Abs(projectPath)
	fmt.Println("Processing path", root)
	ext = extension
	err := filepath.Walk(root, processPath)
	// elapsed := time.Since(start)
	fmt.Println("count of file", Cnt)
	if err != nil {
		fmt.Println("error:", err)
	}
	return Cnt
}

func processPath(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if path != "." {
		if info.IsDir() {
			// fmt.Println("Directory:", path)
		} else {
			// fmt.Println("File:", path)
			if checkExt(info.Name()) {
				// fmt.Println("File:", info.Name())
				Cnt++
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
