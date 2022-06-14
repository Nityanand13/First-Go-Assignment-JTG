package main

import "fmt"
func toRecover() {
    if r := recover(); r != nil {
        fmt.Println("recovered ", r)
    }
}
func divide(a,b int){
	defer toRecover();
	if(b==0){
		panic("You cant divide bcz divisor is zero");
	}
	fmt.Println(a/b);
}

func main(){
	fmt.Print("Enter any two number  a and b for division: ");
	var a,b int;
	fmt.Scan(&a,&b);
	divide(a,b);
}