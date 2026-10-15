package main

import "fmt"

func main() {
	degree := 49.0
	minVal := 5.0
	maxVal := 25.0

	if degree <= maxVal && degree >= minVal {
		fmt.Printf("pass\n")
	} else {
		fmt.Printf("fail\n")
	}

}