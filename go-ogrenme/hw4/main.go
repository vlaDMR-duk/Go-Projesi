package main

import "fmt"


func main(){
	values := []float64{5.0,25.0,4.9,30.0}
	for _, value := range values{
		fmt.Printf("%s %.1f\t",check(value,5.0,25.0),value)
	}
}

func check(value float64, minVal float64, maxVal float64) string {
	if value <= maxVal && value >= minVal {
		return "pass"
	} else {
		return "fail"
	}
}