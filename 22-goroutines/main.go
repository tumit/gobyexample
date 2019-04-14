package main

import "fmt"

func f(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutines")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Scanln()
	fmt.Println("done")
}
