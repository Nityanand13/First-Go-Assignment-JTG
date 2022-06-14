package main
import (
	// "fmt"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Response struct {
	Id  int `json:"ID"`
	Name string `json:"Name"`
	CreatedAt time.Time `json:"CreatedAt"`
	Married bool `json:"Married"`
	LastName *string `json:"LastName"`
	CollegeMarks *float64 `json:"CollegeMarks"`
}

func main(){
	file, err := ioutil.ReadFile("st1.json")

	if err!=nil{
		panic(err);
	}
	response1 := Response{}
 
	err= json.Unmarshal([]byte(file), &response1)
	if err!=nil{
		panic(err);
	}
	fmt.Println(response1.Name);
	fmt.Println(response1.Id);
}