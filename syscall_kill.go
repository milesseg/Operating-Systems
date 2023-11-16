package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
	pid := os.Getpid()
	fmt.Println("Current process PID:", pid)

	err := syscall.Kill(pid, syscall.SIGINT) //TODO: REPLACE SIGKILL, SIGINT
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	time.Sleep(3 * time.Second)
	fmt.Println("Process terminated successfully")
}
