package main

import (
	"fmt"
)

func main(){
	squrt := make(chan int)
	natureal := make(chan int)

	//给natural通道发送数据
	go func(){
		for x:= 0;;x++{
			natureal <- x
		}
	}()

	go func(){
		for {
			x,ok := <-natureal
			if(!ok){
				break;
			}
			squrt <- x*x
		}
	}()

	for {
        fmt.Println(<-squrt)
    }
}