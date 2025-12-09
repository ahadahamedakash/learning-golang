package main

import (
	"bufio"
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

	buf := make([]byte, 11)

	d, err := file.Read(buf)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(buf); i++ {
		fmt.Println("DATA: ", d, string(buf[i]))
	}

	data, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data)) // hellow world

	// fmt.Println("DATA: ", d, buf)

	// READ FOLDERS
	dir, err := os.Open("../")
	if err != nil {
		panic(err)
	}

	defer dir.Close()

	folderInfo, err := dir.ReadDir(-1)

	for _, fi := range folderInfo {
		fmt.Println(fi.Name())
	}

	// CREATE A FILE
	fo, err := os.Create("example_two.txt")

	if err != nil {
		panic(err)
	}

	defer fo.Close()

	fo.WriteString("Write a text!\n")
	fo.WriteString("Write another text!\n")

	bytesTwo := []byte("hello\n")

	fo.Write(bytesTwo)

	// Read and Write to another file
	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	destFile, err := os.Create("example_three.txt")
	if err != nil {
		panic(err)
	}

	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {
		by, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}

			break
		}
		e := writer.WriteByte(by)

		if e != nil {
			panic(e)
		}

	}

	writer.Flush()

	fmt.Println("written to new file successfully!")

	// DELETE A FILE
	fileToDelete, err := os.Open("example_four.txt")
	if err != nil {
		panic(err)
	}

	defer fileToDelete.Close()

	os.Remove(fileToDelete.Name())
}
