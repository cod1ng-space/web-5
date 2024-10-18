package main

import (
	"fmt"
)

// реализовать calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int
func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		select {
		case value := <-firstChan:
			out <- value * value
		case value := <-secondChan:
			out <- value * 3
		case <-stopChan:

		}
	}()
	return out
}

func main() {
	channel1, channel2 := make(chan int), make(chan int)
	stopChan := make(chan struct{})
	calc := calculator(channel1, channel2, stopChan)

	// Тестирование трех случаев
	channel1 <- 5
	//channel2 <- 100
	//stopChan <- struct{}{}
	fmt.Println(<-calc)
}
