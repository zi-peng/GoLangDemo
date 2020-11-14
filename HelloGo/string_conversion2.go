/*
 * @,@Author: ,: your name
 * @,@Date: ,: 2020-11-10 15:48:16
 * @,@LastEditTime: ,: 2020-11-11 10:34:20
 * @,@LastEditors: ,: Please set LastEditors
 * @,@Description: ,: In User Settings Edit
 * @,@FilePath: ,: \GoLangDemo\HelloGo\string_conversion2.go
 */
package main

import (
	"fmt"
	"math"
	"strconv"
)

var orig string = "abc"

var newS string

func conversiondemo() {
	errDemo()
	var orig string = "ABC"
	var newS string
	fmt.Printf("当前ints的大小:%d\n", strconv.IntSize)
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
	}
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	//将int转换为string   %v 也就是value的意思,表示相应值的默认个还是  %s 是转义为字符串的意思
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)

}

/**
 * @,@param {type} argName I am argument argName.
 */
func errDemo() {
	v, ok := mySqrt(25)
	fmt.Println("111", v, "11", v, ok)
	// f, err := os.Open("C:/Go/robots.txt")
	// fmt.Printf("%v", f)
	// if err != nil {
	// 	fmt.Printf("您获取文件报错了，错误是：%v", err)
	// 	return
	// }
	// fmt.Printf("没有报错恭喜")
}

func mySqrt(f float64) (v float64, ok bool) {
	if f < 0 {
		return
	}
	return math.Sqrt(f), true
}

