/*
 * @,@Author: ,: your name
 * @,@Date: ,: 2020-11-11 10:35:28
 * @,@LastEditTime: ,: 2020-11-11 10:36:25
 * @,@LastEditors: ,: Please set LastEditors
 * @,@Description: ,: In User Settings Edit
 * @,@FilePath: ,: \GoLangDemo\HelloGo\season.go
 */
package main

func season(month int) string {
	switch month {
	case 12, 1, 2:
		return "Winter"
	case 3, 4, 5:
		return "Spring"
	case 6, 7, 8:
		return "Summer"
	case 9, 10, 11:
		return "Autumn"
	}
	return "Season unknown"
}
