package task3

import (
	"fmt"
	"sync"
)

var ch = make(chan float64, 4)
var Wg = sync.WaitGroup{}

func ShowResult() {

	Wg.Add(2)

	go CalculateCuboid(7.3 , 5.3 , 6.2)
	go ResultCuboid(ch)

	Wg.Wait()
}

func ResultCuboid(ch <-chan float64) {
	
	for data := range ch {
		fmt.Printf("\nResult : %.2f", data)
	}

	Wg.Done()
}

func CalculateCuboid(Length, Width, Height float64) {

	defer func() {
		close(ch)
		Wg.Done()
	}()

	area := (2 * Length * Width) + (2 * Length * Height) + (2 * Width * Height)
	around := 4 * (Length + Width + Height)
	volume := Length * Width * Height

	ch <- area
	ch <- around
	ch <- volume
}


