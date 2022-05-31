package main

import (
	"fmt"
	"intermedietGO/task2"
	"intermedietGO/task3"
	"sync"
)

type deretBilangan struct {
	limit int
}

var wait = sync.WaitGroup{}

func main()  {

	fmt.Println("Aplikasi Berjalan")

	// Task 1
	deret := deretBilangan{40}

	wait.Add(4)
	go deret.prima()
	go deret.genap()
	go deret.ganjil()
	go deret.fibo()

	wait.Wait()

	// Task 2
	task2.Show()

	// Task 3
	task3.ShowResult()
	
}

func (deret *deretBilangan) prima() {
	result := []int{}
	for i := 1; i < deret.limit; i++ {
		fact := 0
		for j := 1; j < deret.limit; j++ {
			if i%j == 0 {
				fact++
			}
		}
		if fact == 2 && i != 1 {
			result = append(result, i)
		}
	}

	fmt.Printf("Hasil Bil. Prima: %v \n", result)
	wait.Done()
}

func (deret *deretBilangan) genap()  {
	result := []int{}
	for i := 0; i < deret.limit; i++ {
		if i%2 == 0 {
			result = append(result, i)
		}
	}

	fmt.Printf("Hasil Bil. Genap : %v \n", result)
	wait.Done()
}

func (deret *deretBilangan) ganjil()  {
	result := []int{}
	for i := 0; i < deret.limit; i++ {
		if i%2 != 0 {
			result = append(result, i)
		}
	}

	fmt.Printf("Hasil Bil. Ganjil : %v \n", result)
	wait.Done()
}

func (deret *deretBilangan) fibo()  {
	result := []int{}
	bil1, bil2, bil3 := 0, 1, 0

	for i := 1; i < deret.limit; i++ {
		if bil1 < deret.limit {
			result = append(result, bil1)
		}
		bil3 = bil1 + bil2
		bil1 = bil2
		bil2 = bil3
	}
	
	fmt.Printf("Hasil Bil. Fibonacci: %v \n", result)
	wait.Done()
}