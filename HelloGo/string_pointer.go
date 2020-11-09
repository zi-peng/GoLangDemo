/*
 * @,@Author: ,: your name
 * @,@Date: ,: 2020-11-09 20:03:23
 * @,@LastEditTime: ,: 2020-11-09 21:02:31
 * @,@LastEditors: ,: Please set LastEditors
 * @,@Description: ,: In User Settings Edit
 * @,@FilePath: ,: \GoLangDemo\HelloGo\string_pointer.go
 */

package main

import f "fmt"

func stringPointer() {
	// s := "good bye"
	// t := s

	// p := &s
	// //通过*p 通过内存地址更改字符串的值，则该地址也会发生变化，也就是说 s的值会和 *p赋的值一致
	// *p = "ciao"
	// f.Printf("内存地址p:%p\n", p)
	// f.Printf("*p:%s\n", *p)
	// f.Printf("s字符串p:%s\n", s)
	// f.Printf("t字符串p:%s\n", t)
	var p2 *int
	*p2 = 2
	f.Printf("%d", p2)
	//go语言不能获取常量的内存地址，以及字面值的地址
	//const i = 5
	//ptr := &i
	//ptr2 := 10
	//f.Printf("%p", ptr)
	//f.Printf("%p", ptr2)
}
