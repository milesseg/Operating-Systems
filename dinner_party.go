package main

import (
	"fmt"
	"sync"
)

const NUM_PHILOSOPHERS int = 5
const NUM_CHOPSTICKS int = 5
const EAT_TIMES int = 3
const NUM_EATING_PHILOSOPHER int = 2

type Host struct {
	// channel for a philosopher requesting to eat
	requestChannel chan *Philosopher

	// channel for allowing a philosopher to eat
	eatingChannel chan *Philosopher

	// channel that quits the program; tells the host
	// to stop hosting
	quitChannel chan int

	// keep track/bookkeeping of currently eating philosophers
	eatingPhilosophers map[int]bool

	// we need to lock the bookkeeping variable
	mu sync.Mutex
}

func (h *Host) manage() {
	for {

		select {
		// handling a philosopher request
		case p := <-h.requestChannel:
			fmt.Printf("%d submitted request\n", p.ID)

			select {
			// eating channel is not full
			case h.eatingChannel <- p:
			// eating channel is full
			default:
				// finished := <-h.requestChannel // Pops a Philosopher object off the channel
				finished := <-h.eatingChannel
				h.eatingChannel <- p
				currentlyEating := make([]int, 0, NUM_PHILOSOPHERS)
				for index, eating := range h.eatingPhilosophers {
					if eating {
						currentlyEating = append(currentlyEating, index)
					}
				}
				fmt.Printf("%v have been eating, clearing plates %d\n", currentlyEating, finished.ID)

				h.eatingPhilosophers[finished.ID] = false
				h.eatingPhilosophers[p.ID] = true
			}
		default:

		}

		// similar to a switch stmt
		select {
		case <-h.quitChannel:
			// when the channel receives a signal
			// end the host manage function
			fmt.Println("party is over")
			return
		default:
		}
	}

}

func main() {
	fmt.Println("Dinner party commence!!")

	// Declaring a waitgroup to keep track of our threads
	var wg sync.WaitGroup

	// Make our channels for managing requests
	requestChannel := make(chan *Philosopher, NUM_EATING_PHILOSOPHER)
	//eatingChannel :=
	// The quit channel signals the Host to stop managing the dinner party
	// because we're ending the program
	quitChannel := make(chan int, 1)

	// Create a Host
	host := Host{
		requestChannel:     requestChannel,
		quitChannel:        quitChannel,
		eatingPhilosophers: make(map[int]bool),
	}

	// make chopsticks
	chopsticks := make([]*ChopStick, NUM_CHOPSTICKS)
	// GO IS DUMB: it has a range operator,
	// but you can't use it for Integers
	//     for i := range NUM_CHOPSTICKS {
	for i := 0; i < NUM_PHILOSOPHERS; i++ {
		chopsticks[i] = &ChopStick{
			ID: i + 1,
		}
	}

	// make philosophers
	philosophers := make([]*Philosopher, NUM_PHILOSOPHERS)

	for i := 0; i < NUM_PHILOSOPHERS; i++ {
		philosophers[i] = &Philosopher{
			ID:             i + 1,
			Name:           "",
			LeftChopStick:  chopsticks[i],
			RightChopStick: chopsticks[(i+1)%5],
			Host:           &host,
		}
	}

	go host.manage()

	for _, philosopher := range philosophers {
		fmt.Printf("%d philosopher going to eat\n", philosopher.ID)
		go philosopher.Eat(&wg)
	}

	wg.Wait()
	host.quitChannel <- 1

}
