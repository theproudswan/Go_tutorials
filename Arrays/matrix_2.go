/*
1. Сумма элементов двумерного массива
2. Суммы элементов строк матрицы
3. Сумма элементов столбцов матрицы
4. Найти строку матрицы с максимальной суммой элементов
5. Найти индексы максимальных элементов матрицы
6. Максимальные элементы столбцов матрицы
7. Поменять местами строки матрицы
8. Отсортировать в матрице столбцы по убыванию значений элементов в первой строке
9. Найти максимальный элемент диагонали
10. Количество отрицательных элементов под главной диагональю матрицы
11. Найти минимальный элемент матрицы ниже побочной диагонали
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 1. Сумма элементов двумерного массива
func SumMatrix(matrix [][]int) int {
	sum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			sum += matrix[i][j]
		}
	}

	return sum
}

// 2. Суммы элементов строк матрицы
func SumRows(matrix [][]int) []int {
	rows := make([]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		rows[i] = 0
		for j := 0; j < len(matrix[i]); j++ {
			rows[i] += matrix[i][j]
		}
	}

	return rows
}

// 3. Сумма элементов столбцов матрицы
func SumColomns(matrix [][]int) []int {
	rows := make([]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		rows[i] = 0
		for j := 0; j < len(matrix[i]); j++ {
			rows[i] += matrix[j][i]
		}
	}

	return rows
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var size int
	fmt.Print("Введите размер квадратной матрицы: ")
	fmt.Scanf("%d\n", &size)

	var matrix [][]int

	for i := 0; i < size; i++ {
		matrix = append(matrix, make([]int, size))
		for j := 0; j < size; j++ {
			matrix[i][j] = rand.Intn(10-(-10)) + (-10)
		}
	}

	for i := 0; i < size; i++ {
		fmt.Println(matrix[i])
	}

	// Output:
	fmt.Println(SumMatrix(matrix))
	fmt.Println(SumRows(matrix))
	fmt.Println(SumColomns(matrix))
}
