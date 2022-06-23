package main

import "fmt"

func fizzbuzz(num int) {
	for i := 0; i < num; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		}
		fmt.Println(i)
	}
}

func main() {
	fizzbuzz(100)
}
