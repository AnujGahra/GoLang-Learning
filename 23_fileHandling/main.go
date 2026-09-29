package main

import (
	// "fmt"
	"bufio"
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


	// read folders
	// dir, err := os.Open("../")

	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1)

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name())
	// }


	// create file

	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// f.WriteString("Hii, GoLang")
	// f.WriteString(" Proper fileHandling")

	// bytes := []byte("Hello GoLang")
	// f.Write(bytes)


	// read and write to another file (sreaming fashion)

	sourceFile, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	destFile, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}

	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {

				panic(err)
			}
			break
		}

		e := writer.WriteByte(b)
		if err != nil {
			panic(e)
		}


	}

	writer.Flush()
	fmt.Print("Written to new file successfully")


}
