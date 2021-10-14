package main

import "fmt"

func main(){
	var a[3]int //声明
	fmt.Println(a[0]) //打印第一个
	fmt.Println(a[len(a)-1]) //打印末尾

	for i,v := range a{ //打印键和值
		fmt.Printf("%d %d \n",i,v)
	}

	for _,v := range a{ //只打印值
		fmt.Println(v)
	}
}