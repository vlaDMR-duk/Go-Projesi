package main

import (
	"fmt"
	"errors"
)

type Rule struct{
	Feature string
	Min float64
	Max float64
	Critical bool
}


func main(){

	rules := []Rule{
		{Feature: "temperature_mean",Min: 5.0,Max: 25.0,Critical: true},
		{Feature: "precipitation:", Min: 800.0, Max: 300,Critical: false},
		{Feature: "", Min: 6.0, Max: 8.0,Critical: false},
		{Feature: "soil_ph", Min: 6.0, Max: 8.0,Critical: true},
	}

	for _, rule := range rules{
		err := validateRule(rule)
		if err == nil{
			fmt.Println("geçerli: "+ rule.Feature)
		} else {
			fmt.Println(err)
		}
	}
}


func validateRule(rule Rule) error{
	if rule.Feature == "" && rule.Min >= rule.Max{
		return fmt.Errorf("HATA-> %s geçersiz. Feature adi yok, %.1f > %.1f", rule.Feature, rule.Min, rule.Max)
	}else if rule.Feature == ""{
		return errors.New("HATA -> Feature ismi yok")
	} else if rule.Min >= rule.Max{
		return fmt.Errorf("HATA-> %s : geçersiz. Min(%.1f)>=Max(%.1f)", rule.Feature, rule.Min, rule.Max)
	} else {
		return nil
	}
}
func validateRule2(rule Rule) error {
	if rule.Feature == "" {
		return errors.New("feature adi bos")
	}																																				//validaterule un daha kısa yazılışı
	if rule.Min > rule.Max {
		return fmt.Errorf("kural %s gecersiz: min (%.1f) > max (%.1f)", rule.Feature, rule.Min, rule.Max)
	}
	return nil
}