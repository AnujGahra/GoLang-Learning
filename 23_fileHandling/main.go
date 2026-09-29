package main

import (
	"fmt"
	"os"
)


func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	// log the error 
	// 	panic(err)
	// }
	// fmt.Println("File Name: ", fileInfo.Name())
	// fmt.Println("File IsDir: ", fileInfo.IsDir())
	// fmt.Println("File Size: ", fileInfo.Size())
	// fmt.Println("File Perission: ", fileInfo.Mode())
	// fmt.Println("File modified at: ", fileInfo.ModTime())
	

	// read file
// 	f, err := os.Open("example.txt")
// 	if err != nil {
// 		panic(err)
// 	}

//  defer f.Close()

//  buf := make([]byte, 10)
//  d, err := f.Read(buf)
//  if err != nil {
// 	panic(err)
//  }

//  fmt.Println("data", d,  buf)

//  for i := 0; i < len(buf); i++ {
// 	fmt.Println("data", d, string(buf[i]))
//  }

	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))


	

}
