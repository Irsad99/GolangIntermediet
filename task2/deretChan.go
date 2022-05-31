package task2

import (
	"fmt"
	"sync"
)

var Wg = sync.WaitGroup{}
var ch = make(chan []int)

func Show() {

	Wg.Add(2)
	go GanjilGenap(ch)
	go Fibonacci(40, ch)

	Wg.Wait()
}

func GanjilGenap(ch <-chan []int) {
	getData := <-ch
	resultGanjil := []int{}
	resultGenap := []int{}
	for i := 0; i < len(getData); i++ {
		if getData[i]%2 == 0 {
			resultGenap = append(resultGenap, getData[i])
		} else {
			resultGanjil = append(resultGanjil, getData[i])
		}
	}

	fmt.Printf("\nHasil Bil. Genap : %v \n", resultGenap)
	fmt.Printf("Hasil Bil. Ganjil : %v \n", resultGanjil)
	Wg.Done()
}

func Fibonacci(data int, ch chan<- []int) {
	result := []int{}
	bil1, bil2, bil3 := 0, 1, 0

	for i := 1; i < data; i++ {
		if bil1 < data {
			result = append(result, bil1)
		}
		bil3 = bil1 + bil2
		bil1 = bil2
		bil2 = bil3
	}

	ch <- result
	Wg.Done()
}
