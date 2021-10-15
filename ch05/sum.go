package main

import "fmt"

func main()  {
	fmt.Println(sum(1,2,3))
	fmt.Println(sum())
}

func sum(args...int) int {	
	if(len(args) == 0){
		return 0
	}

	sum := 0

	for _,arg := range args {
		sum += arg
	}
	return sum
}