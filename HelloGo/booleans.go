/*
 * @,@Author: ,: your name
 * @,@Date: ,: 2020-11-10 12:24:13
 * @,@LastEditTime: ,: 2020-11-10 18:25:24
 * @,@LastEditors: ,: Please set LastEditors
 * @,@Description: ,: In User Settings Edit
 * @,@FilePath: ,: \GoLangDemo\HelloGo\booleans.go
 */
package main

import (
	"fmt"
	"runtime"
)

func init() {

	if runtime.GOOS == "windows" {
		fmt.Printf(runtime.GOOS)
	}
}

func demo2() {
	if str1 := ""; str1 == "222" {
		fmt.Printf("str1的值是%s", str1)
	}
	//fmt.Printf("str1括号外的值是%s", str1)
}
func isGreater(x, y int) bool {
	if x > y {
		return true
	}
	return false
}
func demo(x int) int {
	if x < 10 {
		return 1
	} else if x > 10 {
		return 0
	}
	return 0
}
