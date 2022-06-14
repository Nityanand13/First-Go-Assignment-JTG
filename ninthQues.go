package main

import (
	"fmt"
	"os"
	"crypto/md5"
)
func main(){
	file, err := os.Open("./file.txt")
	defer file.Close();
	if err != nil {
		panic(err);
	}
	hashCode := md5.New();
	fmt.Printf("%x", hashCode.Sum(nil));
}
