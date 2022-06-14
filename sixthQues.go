package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	file1, err := os.Create("./source.txt")

	if err != nil {
		panic(err)
	}
	fileContent := "Hii this is my source file!"
	length1, err := io.WriteString(file1, fileContent)
	if err != nil {
		panic(err);
	}
	fmt.Println(length1);
	contentOfSourceFile,err := ioutil.ReadFile("./source.txt");
	if err != nil {
		panic(err);
	}
	file2, err := os.Create("./dest.txt")
	if err != nil {
		panic(err)
	}
	length2, err := io.WriteString(file2, string(contentOfSourceFile));
	if err != nil {
		panic(err)
	}
	fmt.Println(length2);
	contentOfDestFile,err := ioutil.ReadFile("./dest.txt");
	if err != nil {
		panic(err);
	}
	fmt.Println(string(contentOfDestFile));
	defer file1.Close();
	defer file2.Close();
}
