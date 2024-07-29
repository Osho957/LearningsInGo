package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Working with files")
	content := "put these content in file"
	file, err :=os.Create("./file.txt")
	if err != nil {
		panic(err)
	}
	length,err := io.WriteString(file,content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is: ",length)
	defer file.Close()
	readFile("./file.txt")
}

func readFile(fileName string){
     databyte, err := os.ReadFile(fileName)
	 if err != nil {
		panic(err)
	}
	fmt.Println("Text data inside the file is \n", string(databyte))
}