//修改echo程序，使其能够打印os.Args[0]，即被执行命令本身的名字。
package main 

import (
	"fmt"
	"os"
)

func main(){
	var s,sep string
	for i:= 0;i < len(os.Args) ; i++{
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
//执行 go run .\test01.go 1  2 3  4
//输出C:\Users\ADMINI~1\AppData\Local\Temp\go-build574482948\b001\exe\test01.exe 1 2 3 4