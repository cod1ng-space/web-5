package main

import "fmt"

// реализовать removeDuplicates(in, out chan string)
func removeDuplicates(inputStream chan string, outputStream chan string) {
	last := ""
	for {
		s, opened := <-inputStream
		if !opened {
			break
		}
		if last != s {
			outputStream <- s
			last = s
		}
	}
	close(outputStream)
}

func main() {
	// здесь должен быть код для проверки правильности работы функции removeDuplicates(in, out chan string)
	inputStream := make(chan string)
	outputStream := make(chan string)
	go removeDuplicates(inputStream, outputStream)
	s := "1112244115668" // Значение для тестирования
	go func() {
		defer close(inputStream)
		for _, symbol := range s {
			inputStream <- string(symbol)
		}
	}()

	for symbol := range outputStream {
		fmt.Print(symbol)
	}
}
