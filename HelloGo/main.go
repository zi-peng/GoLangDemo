/*
 * @Author: your name
 * @Date: 2020-11-09 17:26:01
 * @,@LastEditTime: ,: 2020-11-09 19:11:02
 * @,@LastEditors: ,: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \goCode\main.go
 */
package main // 声明 main 包，表明当前是一个可执行程序
import (
	f "fmt"
)

const b = 5

// 导入内置 fmt 包

func main() { // main函数，是程序执行的入口
	f.Println("hello,12311131")
}

func funcDemo(a int, b int) (int, int) {
	f.Printf("%b", a+b)
	return 0, a
}
