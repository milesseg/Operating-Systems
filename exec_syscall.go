package main

import (
	"fmt"
	"syscall"
)

func main() {
	err := syscall.Exec("/bin/ls", []string{"ls", "-la"}, syscall.Environ())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Process created successfully:")
}
