package main

import "fmt"

func main() {
	degree := 17.4
	minVal := 5.0
	maxVal := 25.0

	if degree <= maxVal && degree >= minVal {
		fmt.Printf("Pass \n")
	} else {
		fmt.Printf("fail")
	}

}