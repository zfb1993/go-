package main

import "fmt"

func main(){
	ages := make(map[string]int)
	ages["you"] = 888
	ages["my"] = 18

	name := map[string]string{
		"your":"ss",
		"my":"bb",
	}

	delete(name,"your")
	fmt.Println(name)
	fmt.Println(ages)
}