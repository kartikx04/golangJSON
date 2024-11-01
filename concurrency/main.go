package main

import (
	"fmt"
	"sync"
	"time"
)

func generateNumbers(total int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= total; i++ {
		fmt.Printf("Generated Number: %d\n", i)
		time.Sleep(10 * time.Millisecond)
	}
}

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 20; i++ {
		fmt.Printf("Printed Number: %d\n", i)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go printNumbers(&wg)
	go generateNumbers(20, &wg)

	wg.Wait()
}
