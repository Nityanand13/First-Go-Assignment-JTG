package main

import (
	// "fmt"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)
type Response struct {
	Id  int
	Name string
	CreatedAt time.Time
	Married bool
	LastName *string
	CollegeMarks *float64
}


func main(){
	var lname string="Rai";
	var marks float64=93;
	
	// facing problem in this time
	t, err := time.Parse(time.UnixDate, "Mon Feb 14 11:06:39 IST 2022");
	if err != nil {
		panic(err)
	}
	Response1 :=&Response{Id: 1,Name: "Shivam",CreatedAt:t,Married:  false,LastName: &(lname),
	CollegeMarks: &(marks)};
	file, err := json.MarshalIndent(Response1,"","");
	if err!=nil{
		fmt.Print(err);
	}
	ioutil.WriteFile("st1.json", file,0644);
}

// Dj9e@rT639tjAmH