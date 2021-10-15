package main

import "fmt"

func main()  {
	fmt.Println(f())
	fmt.Println(f2())
}

func f() (result int) {
    defer func() {
        result++
    }()
    return 0
}

func f2() (r int) {
	t := 5
	defer func() {
	  t = t + 5
	}()
	return t
}