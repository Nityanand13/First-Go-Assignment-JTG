package main

import "fmt"

func sum(slice []int)(int,string) {
	var allSum int=0;
	for i:=0;i<len(slice);i++{
		allSum+=slice[i];
	}
	if allSum%2==0{
		return allSum,"true"
	} else{
		return allSum,"false"
	}
}
func main(){
	slice:=[]int{1,5,9,4};
	var a,b= sum(slice);
	fmt.Printf("%d %s",a,b);
	fmt.Println();
	// fmt.Println(b);
	// fmt.Println(a+" "+b);
}