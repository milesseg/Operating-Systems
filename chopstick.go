package main

import "sync"

type ChopStick struct {
	sync.Mutex

	ID int
}
