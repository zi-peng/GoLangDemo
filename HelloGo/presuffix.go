/*
 * @,@Author: ,: your name
 * @,@Date: ,: 2020-11-09 16:20:03
 * @,@LastEditTime: ,: 2020-11-09 18:18:38
 * @,@LastEditors: ,: Please set LastEditors
 * @,@Description: ,: In User Settings Edit
 * @,@FilePath: ,: \undefinede:\goCode\H\presuffix.go
 */
package main // 声明 main 包，表明当前是一个可执行程序

import (
	"fmt"

	"strings"
)

func main1() {
	var str string = "This is an example of a string"
	// \表示转义符，需要转义的字符放在该符号的的后面就行，  %s 表示字符串占位符,直接写在输出字符串中即可， 最后的 Th 第二个站位符(%s)的位置
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s?", str, "Th")
	//判断该语句是否以  Ts开头   %t 表示布尔站位符
	fmt.Printf("%t\n", strings.HasPrefix(str, "Ts"))
	//判断该语句是否以  ng结尾   %t 表示布尔站位符
	fmt.Printf("%t\n", strings.HasSuffix(str, "ng"))
}
