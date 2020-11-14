/*
 * @,@Author: ,: your name
 * @,@Date: ,: 2020-11-11 15:15:32
 * @,@LastEditTime: ,: 2020-11-11 16:25:53
 * @,@LastEditors: ,: Please set LastEditors
 * @,@Description: ,: In User Settings Edit
 * @,@FilePath: ,: \GoLangDemo\HelloGo\for1.go
 */
package main

import "fmt"

func for1Demo() {
	for i := 0; i < 5; i++ {
		fmt.Printf("输出%d\n", i)
	}
}

func rangeDemo() {
	for i := 0; i < 3; {
		fmt.Printf("该值是:%d", i)
	}
}
