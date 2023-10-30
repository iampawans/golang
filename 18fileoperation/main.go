package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to File Operation learning in GO...")

	filepath := "./MyFile.txt"
	content := "Hi There, Welcome to the learning of file operation in GO"
	file := CreateFile(filepath)
	WriteInFile(file, content)
	ReadFile(filepath)
}

func CreateFile(filename string) *os.File {
	file, err := os.Create(filename)
	CheckNilErr(err)
	return file
}

func WriteInFile(file *os.File, content string) {
	length, err := io.WriteString(file, content)
	CheckNilErr(err)
	fmt.Println("Size of text written to the file is :", length)
}

func ReadFile(filename string) {
	//ioutil.ReadFile() is deprecated
	//Return type is []byte
	databyte, err := os.ReadFile(filename)
	CheckNilErr(err)
	fmt.Println("databyte : ", databyte)
	content := string(databyte)
	fmt.Println("Content of File : ", content)
}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
