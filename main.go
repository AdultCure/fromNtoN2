package main

import (
	"errors"
	"fmt"
)

func main() {
	printFromTo(10)
}

func printFromTo(start int) {
	end := start * start

	array, err := calculateFromTo(start, end)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Начало: %d, конец: %d\n", start, end)
	fmt.Printf("%d", array)
}

func calculateFromTo (start, end int) ([]int, error){
	if end == 0 {
		return make([]int, 0), errors.New("Не может равняться нулю")
	}

	var array []int

	for i := start; i <= end; i++ {
		fmt.Printf("%d", i)
		array = append(array, i)
	}

	return array, nil
}
