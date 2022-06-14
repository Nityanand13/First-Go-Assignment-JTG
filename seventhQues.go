package main

import (
	"fmt"
	"time"
)

func fun() {
	time.Sleep(5*time.Millisecond)
    fmt.Println("Hello world");
}

func main(){
	go fun();
	go fun();
	go fun();
	go fun();
	go fun();
	time.Sleep(time.Second);
}