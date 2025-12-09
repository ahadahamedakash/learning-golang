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

	fmt.Println(fileInfo.Name())
	// fmt.Println(fileInfo.IsDir())
	// fmt.Println(fileInfo.Size())
	// fmt.Println(fileInfo.Mode())
	// fmt.Println(fileInfo.ModTime())

	file, err := os.Open("example.txt")
	if err != nil {
		// log the error
		panic(err)
	}

	defer file.Close()

	buf := make([]byte, 20)

	d, err := file.Read(buf)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(buf); i++ {
		fmt.Println("DATA: ", d, string(buf[i]))
	}

	// fmt.Println("DATA: ", d, buf)

}
