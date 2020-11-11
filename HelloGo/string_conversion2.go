/*
 * @,@Author: ,: your name
 * @,@Date: ,: 2020-11-10 15:48:16
 * @,@LastEditTime: ,: 2020-11-10 18:26:36
 * @,@LastEditors: ,: Please set LastEditors
 * @,@Description: ,: In User Settings Edit
 * @,@FilePath: ,: \GoLangDemo\HelloGo\string_conversion2.go
 */
package main

import (
	"fmt"
	"strconv"
)

var orig string = "abc"

var newS string

func conversiondemo() {

	var orig string = "ABC"
	var newS string
	fmt.Printf("当前ints的大小:%d\n", strconv.IntSize)
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
	}
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}
