/*
Дан массив из 50 элементов, значения которых формируются функцией random и лежат в диапазоне от -50 до 49 включительно.
Требуется из одного массива скопировать в другой массив значения в диапазоне от -5 до 5 включительно и подсчитать их количество.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var arr [50]int
	var result []int

	for i := 0; i < len(arr); i++ {

		arr[i] = rand.Intn(49-(-50)) + (-50)

		if -5 <= arr[i] && arr[i] <= 5 {
			result = append(result, arr[i])
		}

	}

	fmt.Println(arr)
	fmt.Println()
	fmt.Println("Всего элементов в диапазоне [-5..5]: ", len(result))
	fmt.Println(result)
}
