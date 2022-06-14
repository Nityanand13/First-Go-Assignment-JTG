package main

import "fmt"

type shape interface{
	perimeter();
}
type rectangle struct{
	length, breadth float64;
}

type circle struct{
	radius float64
}

func (r rectangle)perimeter() float64 {
    return 2*r.length+2*r.breadth;
}
func (c circle)perimeter() float64 {
    return 2*3.142*c.radius;
}

func main(){
	var length,breadth,radius float64;
	fmt.Println("Enter length and breadth of rectangle: ");
	fmt.Scan(&length,&breadth);
	r := rectangle{length,breadth};
	fmt.Println("Perimeter of rectangle is: ",r.perimeter());
	fmt.Println("Enter radius of circle: ");
	fmt.Scan(&radius);
    c := circle{radius};
	fmt.Println("Perimeter of circle is: ",c.perimeter());
}