package main

import "fmt"

/*go-plus gives an error, because there is another func main in hello2.go. Running
+this code by itself with "go run multiplication.go" will be fine.
*/

func main() {
	fmt.Println(32132 * 42452)
}
