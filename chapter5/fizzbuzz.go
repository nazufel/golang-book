package main

import "fmt"

func main() {
	i := 0
	/* Willingly ignoring switches and else statements, according to advice from
	   "Just for Func" https://www.youtube.com/watch?v=4D506W1AjeM&index=30&list=PL64wiCrrxh4Jisi7OcCJIUpguV_f5jGnZ
	*/
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
