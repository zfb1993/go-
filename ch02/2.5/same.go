// 判断两个字符串是否是相互打乱的
package main

import "fmt"

func main() {
	str1 := "qwe"
	str2 := "ewq"

	str1Map := make(map[rune]int)
	str2Map := make(map[rune]int)

	for _,i := range str1 {
		str1Map[i]++
	}

	for _,i := range str2 {
		str2Map[i]++
	}

	for i,v := range str1Map { 
		if str2Map[i] != v {
			fmt.Println(false)
		}
	}
	fmt.Println(true)
}