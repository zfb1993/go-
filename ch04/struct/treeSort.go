package main

import "fmt"

type Tree struct{
	age int
	heigh int
}

func main(){
	a := 4
	t := Tree{10,20}
	printT(&t)
	printC(&a)
}

func printT(t *Tree){
	fmt.Println(t)
}

func printC(a *int){
	fmt.Println(a)
}