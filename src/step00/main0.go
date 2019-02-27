package main

import (
	"fmt"
    "os"
    "path/filepath"	
)
func currentDir() {
    // Current directory
    ex, err := os.Executable()
    if err != nil {
        panic(err)
    }
    exPath := filepath.Dir(ex)
    fmt.Println("Current directory : ")
    fmt.Println(exPath)	
}

func absoluteFilePath(fileName string) {
    fmt.Println("Absolute file path for '" + fileName + "' :")
    s, _ := filepath.Abs(fileName)
	fmt.Println(s)
}

func main() {
    fmt.Println("hello world")
    fmt.Println("")
    currentDir()
    fmt.Println("")
    absoluteFilePath("foo.txt")
}
