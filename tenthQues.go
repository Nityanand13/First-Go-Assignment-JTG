package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {
	fileContent,err := ioutil.ReadFile("./file.txt");
	if err != nil {
		panic(err);
	}
    encodedStr := base64.StdEncoding.EncodeToString(fileContent);
    fmt.Println("Encoded:", encodedStr);
    decodedStr, err := base64.StdEncoding.DecodeString(encodedStr);
    if err != nil {
        panic("malformed input");
    }
    fmt.Println("Decoded:", string(decodedStr));
}