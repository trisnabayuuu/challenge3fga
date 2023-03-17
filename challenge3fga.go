package main

import "fmt"

func main() {
	data := "hello world"

	// i := range data
	for _, i := range data {
		fmt.Printf("%c \n", i)
	}

	count := make(map[string]int)

	for _, i := range data {
		count[string(i)] += 1

	}
	fmt.Println(count)
}

///////
