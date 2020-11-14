/*
 * @,@Author: ,: your name
 * @,@Date: ,: 2020-11-11 20:48:25
 * @,@LastEditTime: ,: 2020-11-11 20:52:11
 * @,@LastEditors: ,: Please set LastEditors
 * @,@Description: ,: In User Settings Edit
 * @,@FilePath: ,: \GoLangDemo\HelloGo\for6.go
 */
package main

import "fmt"

func for6Demo() {
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i的值是%d,j的值是%d\n", i, j)
		}
	}
}
