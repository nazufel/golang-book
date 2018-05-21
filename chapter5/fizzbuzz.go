package main

import "fmt"

func main() {
	i := 0
	for i <= 100 {
		if i%3 == 0 {
			fmt.Println(i, ": is divisible by three, therefore, 'Fizz.'")
		}
		if i%5 == 0 {
			fmt.Println(i, ": is divisible by five, therefore, 'Buzz.'")
		}
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, ": is divisible by three and five, therefore, 'FizzBuzz.'")
		}
		i++
	}
}
