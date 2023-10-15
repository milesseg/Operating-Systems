package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func fetchWiki(city string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	url := fmt.Sprintf("https://www.youtube.com/%s", city)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error fetching page for %s: %s", city, err)
	}

	//fmt.Println(city)

	defer resp.Body.Close()

	ch <- fmt.Sprintf("City: %s", city)
}

func main() {
	start := time.Now()

	cities := []string{"Manila", "Seoul", "Paris", "Detroit"}

	ch := make(chan string)
	var wg sync.WaitGroup

	for _, city := range cities {
		wg.Add(1)
		go fetchWiki(city, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)

	}()

	for result := range ch {
		fmt.Println(result)
	}

	fmt.Println("Time: ", time.Since(start))

}
