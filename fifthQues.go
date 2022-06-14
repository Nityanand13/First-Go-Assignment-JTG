package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Create("./file.txt")

	if err != nil {
		panic(err)
	}
	fileContent := "Hii this is my file!"
	length, err := io.WriteString(file, fileContent)
	if err != nil {
		panic(err);
	}
	fmt.Println(length);
	content,err := ioutil.ReadFile("./file.txt");
	if err != nil {
		panic(err);
	}
	fmt.Println(string(content));
}
