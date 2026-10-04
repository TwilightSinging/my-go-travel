package main
import (
	"fmt"
	"calculator/cal"
)
func main(){
	fmt.Println(cal.Add(1,2))
	fmt.Println(cal.Sub(1,2))
	fmt.Println(cal.Mul(1,2))
	fmt.Println(cal.Div(1,2))
	fmt.Println(cal.Div(1,0))
}