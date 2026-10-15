package main

import "fmt"


type Rule struct{
	Feature string
	Min float64
	Max float64
	Critical bool
}

func main(){
	features := map[string]float64{
		"temperature_mean": 17.4,			
		"precipitation": 250.0,
	}

	rules := []Rule{
		{Feature: "temperature_mean", Min: 5.0, Max: 25.0, Critical: true}, 	//slice yapısı
		{Feature: "precipitation", Min: 300.0, Max: 800.0, Critical: false},
		{Feature: "soil_ph", Min: 6.0, Max: 8.0, Critical: false},
	}

	kritikEksik := false

	for _,rule := range rules{
		value, ok := features[rule.Feature]
		if !ok{
			fmt.Println(rule.Feature + " -> unknown")
			if rule.Critical == true{
				kritikEksik = true
			}
		} else {
			sonuc := checkRule(rule,value)
			fmt.Println(rule.Feature + "->" + sonuc)
		}
	}

	if kritikEksik == true{
		fmt.Println("Sonuc:"+" "+"insufficient_data")
	} else{
		fmt.Println("Sonuc:"+" "+"reliable_data")
	}
}

func checkRule(rule Rule, value float64) string{ 
	if value>=rule.Min && value<=rule.Max{
		return "pass"
	} else {
		return "fail"
	}
}
