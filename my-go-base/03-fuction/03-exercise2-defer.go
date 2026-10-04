package main
import "fmt"
func f()(result int){
	defer func(){
		result++
	}()
	//defer 会在函数返回前执行
	//最后面的()是defer的参数，会在函数返回前执行
	//defer func(){}()是匿名函数，会在函数返回前执行
	return 0
}
func main(){
	fmt.Println(f())
}