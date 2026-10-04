package main

import "fmt"

func main() {
	s := make([]int, 3, 5)
	s1 := s[1:3]
	s1 = append(s1, 100)
	//fmt.Println(s[4]) // 答：输出[0,0,0,100,0]
	fmt.Println(s1) // 答：输出[0,0,100]
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	//或者用%v %d来打印
	fmt.Printf("%v len:%d cap:%d\n ", s, len(s), cap(s))
	fmt.Printf("%v len:%d cap:%d\n ", s1, len(s1), cap(s1))
	//展示扩容
	s1 = append(s1, 200)
	s1 = append(s1, 300)
	s1 = append(s1, 400)
	fmt.Printf("value of s1:%v len:%d cap:%d\n ", s1, len(s1), cap(s1))
	fmt.Printf("value of s:%v len:%d cap:%d\n ", s, len(s), cap(s))
	//再改变s1的值不会影响s的值
	s1[0] = 1000
	fmt.Println("==========s1[0] = 1000==========")
	fmt.Printf("value of s1:%v len:%d cap:%d\n ", s1, len(s1), cap(s1))
	fmt.Printf("value of s:%v len:%d cap:%d\n ", s, len(s), cap(s))
}
