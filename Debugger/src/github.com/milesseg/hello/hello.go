package main

import "fmt"

func f() int {
	var n int
	n = 1
	n = 2
	return n
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(f())
}
