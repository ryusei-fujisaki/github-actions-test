package main

import "fmt"

func main() {
	fmt.Println("hell world")
	fmt.Println("hello world fujisaki")
	num := sub(1)
	fmt.Println(num)
	str := sub2()
	fmt.Println(str)
}

func sub(a int) int {
	a += 1
	return a
}

func sub2() string {
	b := "DMM-fujisaki"
	return b
}
