/*
 * @,@Author: ,: your name
 * @,@Date: ,: 2020-11-09 19:48:15
 * @,@LastEditTime: ,: 2020-11-09 20:01:55
 * @,@LastEditors: ,: Please set LastEditors
 * @,@Description: ,: In User Settings Edit
 * @,@FilePath: ,: \GoLangDemo\HelloGo\pointer.go
 */
package main

import "fmt"

//指针-- 获取该变量的存储位置
func pointerDemo() {
	var i1 = 5
	var i2 = 6
	//%p是获取该变量的内存地址占位符，%d是值类型的占位符
	fmt.Printf("该数字是5，%d内存的地址是:%p\n", i1, &i1)
	//初始化一个int型指针
	var intP *int
	intP = &i2
	var p *string
	fmt.Printf("%p,int类型\n", p)
	fmt.Printf("%p,string类型", intP)
}
