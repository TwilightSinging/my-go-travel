// package main

// import (
// 	"fmt"
// 	"unsafe"
// )

// /*
// 给出以下两个结构体，不运行代码，计算它们在 64 位系统上的 unsafe.Sizeof，并解释差异：
// type A struct { a bool; b int64; c int32 }
// type B struct { b int64; c int32; a bool }
// 💡 提示：画出每个字段的内存偏移量和 padding。
// */
// type A struct {
// 	a bool  //1字节
// 	b int64 //8字节
// 	c int32 //4字节
// }

// // 内存对齐
// type B struct {
// 	b int64 //8字节
// 	c int32 //4字节
// 	a bool  //1字节
// }

// //A和B的内存对齐方式不同，A是按字段顺序对齐，B是按字段类型对齐
// //输出A:24 B:16
// /*
// int64的偏移要是8的倍数，bool补齐到8，int32起始偏移是4的倍数，不用补，
// 但最后struct要补全到8的倍数1+7+8+4=20再+4到8的倍数
// A: 1 + 7 + 8 + 4 + 4(补齐) = 24
// B: 8 + 4 + 1 + 3(补齐) = 16*/
// /*
// CPU 读取内存不是一个字节一个字节读的，
// 而是按“字长”（64位系统是 8 字节）一块一块读的。
// 如果 int64 的起始地址不是 8 的倍数，它就会横跨两个内存块，
// CPU 需要读两次并拼接，性能会下降。*/
//
//	func main() {
//		fmt.Println("A:", unsafe.Sizeof(A{}))
//		fmt.Println("B:", unsafe.Sizeof(B{}))
//	}
package main

import (
	"fmt"
	"unsafe"
)

// 对应 C++ 的 struct test2
type Test2 struct {
	A int32   // 对应 int a (4字节)
	C float64 // 对应 double c (8字节)
}

// 对应 C++ 的 struct testMemory
type TestMemory struct {
	A int32 // 对应 int a (4字节)
	B int64 // 对应 long b (8字节)
	C byte  // 对应 char c (1字节)
	L Test2 // 对应 struct test2 l (16字节)
}

func main() {
	fmt.Println("test2 大小:", unsafe.Sizeof(Test2{}))
	fmt.Println("testMemory 大小:", unsafe.Sizeof(TestMemory{}))
}
