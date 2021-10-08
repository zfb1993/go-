//打印命令行参数
package main

import (
	"fmt"
	"os"
)

func main(){
	var s, sep string
    for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	} 
	fmt.Println(s)
}
//输入  1   2  334           4
//输出1 2 334 4 把每个参数用空格分隔了一下