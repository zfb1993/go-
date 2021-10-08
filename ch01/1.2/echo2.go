//打印命令行参数
package main

import (
	"fmt"
	"os"
)

func main(){
	var s,sep string 
	for _,arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
//输入  1  2  3  444    1
//输出1 2 3 444 1 默认的第一个参数是文件本身的名字