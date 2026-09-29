package main

import (
	"fmt"
	"os"
)


func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		// log the error
		panic(err)
	}

	fileInfo, err := f.Stat()
	if err != nil {
		// log the error 
		panic(err)
	}
	// fmt.Println("File Name: ", fileInfo.Name())
	// fmt.Println("File IsDir: ", fileInfo.IsDir())
	// fmt.Println("File Size: ", fileInfo.Size())
	// fmt.Println("File Perission: ", fileInfo.Mode())
	// fmt.Println("File modified at: ", fileInfo.ModTime())

}
