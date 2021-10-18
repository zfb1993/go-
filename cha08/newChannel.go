package main

import (
	"fmt"
)

func main()  {
	squrt := make(chan int)
	nature := make(chan int)

	go func ()  {
		for i := 1; i <= 10; i++{
			nature <- i
		}
		close(nature)
	}()

	go func ()  {
		for i := range nature {
			squrt <- i*i
		}
		close(squrt)
	}()

	for x := range squrt {
		fmt.Println(x)
	}
}