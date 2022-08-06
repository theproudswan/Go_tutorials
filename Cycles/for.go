// Поиск совершенных чисел от 1 до N

package main

import (
	"fmt"
	"os"
)

func IsSimple(num int) bool {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	N := 140000000000
	var prev = 0
	var i = 1
	var tmp int
	for {
		if IsSimple(i + 1) {
			tmp = (1 << (i + 1)) - 1
			if IsSimple(tmp) {
				tmp = tmp << i
				if tmp > N || tmp < prev {
					os.Exit(0)
				}
				fmt.Println(tmp)
				prev = tmp
			}
		}
		i++
	}
}
