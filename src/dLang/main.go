package main

import (
	"fmt"
	"path/filepath"
	"time"
	"dLang/rules/Rule_1"
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
	fmt.Println("The java extension count is %v",rule_1.FindCount(ext, root))
	elapsed := time.Since(start)
	fmt.Println("Source code took %s", elapsed)

}
