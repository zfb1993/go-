//修改echo程序，使其打印每个参数的索引和值，每个一行。
package main

import (
	"fmt"
	"os"
)

func main(){
	for i ,arg := range os.Args[1:] {
		fmt.Printf("索引：%d,值：%s \n",i,arg)
	}
}