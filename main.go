package main

import "fmt"

func ShowFizzBuzz (numbers []int) {

	for count, _ := range numbers {
		
		count += 1

		if count%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if count%3 == 0 {
			fmt.Println("Fizz")
		} else if count%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(count)
		}
	}
}

func main () {	
	
	numbers := make([]int, 100)
	
	ShowFizzBuzz(numbers)
	
}
