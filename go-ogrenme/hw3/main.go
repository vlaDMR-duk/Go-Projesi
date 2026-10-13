package main

import "fmt"

func main(){
	value := 45.0
	minVal := 25.0
	maxVal := 100.0
	sonuc := check(value,minVal,maxVal)
	fmt.Printf("Sonuç = %s\n",sonuc)
}

func check(value float64, minVal float64, maxVal float64) string {
	if value <= maxVal && value >= minVal {
		return "pass"
	} else {
		return "fail"
	}

}