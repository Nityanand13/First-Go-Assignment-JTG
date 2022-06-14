package main
import "fmt"
func main(){
	var str string;
	fmt.Print("Enter an String For verticle posting: ");
	fmt.Scan(&str);
	var length int= len(str);
	for i:=0;i<len(str);i++{
		for j:=0;j<length;j++{
			if i==0{
				fmt.Printf("%c ",str[j]);
			} else { 
				if j==0||j==length-1{
					fmt.Printf("%c",str[i]);
				} else{
					fmt.Printf(" ");
				}
			}
		}
		fmt.Println();
		if i==0{
			length=1;
		}
		length+=2;
	}
}