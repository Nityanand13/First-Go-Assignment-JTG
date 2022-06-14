package main
import (
    "fmt"
    "time"
)

func fun(str string) {
    for i := 1;i<=5;i++ {
        time.Sleep(5*time.Millisecond)
    	fmt.Println(str);
    }
}

func main() {
    go fun("ping");
    go fun("pong");
    time.Sleep(time.Second);
}