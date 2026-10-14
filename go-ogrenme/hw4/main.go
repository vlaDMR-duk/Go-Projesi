package main

import "fmt"


func main(){
	values := []float64{5.0,25.0,4.9,30.0}
	pass := 0
	fail := 0
	for _, value := range values{
		sonuc := check(value,5.0,25.0)
		if sonuc == "pass"{
			pass +=1
		} else {
			fail +=1
		}
	}
	fmt.Printf("pass = %d,fail = %d \n", pass, fail)
}

func check(value float64, minVal float64, maxVal float64) string {
	if value <= maxVal && value >= minVal {
		return "pass"
	} else {
		return "fail"
	}
}