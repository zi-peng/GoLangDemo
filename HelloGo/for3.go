/*
 * @,@Author: ,: your name
 * @,@Date: ,: 2020-11-11 16:28:48
 * @,@LastEditTime: ,: 2020-11-11 20:44:35
 * @,@LastEditors: ,: Please set LastEditors
 * @,@Description: ,: In User Settings Edit
 * @,@FilePath: ,: \GoLangDemo\HelloGo\for3.go
 */
package main

import "fmt"

func for3Demo() {
	var i = 0
	for {
		i = i - 1
		fmt.Printf("输出值%d", i)
		if i < 0 {
			break
		}
	}
}

func for3Demo1() {
	for i := 0; i < 3; i++ {
		for j := 0; i < 10; j++ {
			if j > 5 {
				break
			}
			fmt.Println("j的值：", j)
		}
		fmt.Println("i的值:", i)
	}
}
