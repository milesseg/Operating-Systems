package main

import (
	"time"
)

var philo = []string{"Aristotle", "Socrates", "Plato", "Confucius", "Tzu"}

//Have each philosopher eat for a certain period
//Time how long each phil has waited
//One who has waited longest shall eat next

const eatTime = time.Second / 100
