package main


import "fmt"


type Rule struct{
	Feature string
	Min float64
	Max float64
	Critical bool
}


func main() {
	mission1()
	mission2()
}

func mission1(){
	fmt.Println("Hello world")
	r := Rule{
		Feature: "temperature_mean",
		Min: 5.0,
		Max: 25.0,
		Critical: true,
	}

	fmt.Printf("%s \n",r.Feature)
	sonuc := checkRule(r, 15.0)
	fmt.Println(sonuc)
	
}

func checkRule(rule Rule, value float64) string{ // fonksiyona verirken önce verimiz sonra veri tipimiz koyuluyor, rule bizim verimiz ve veri tipi de oluşturduğumuz Rule struct ı. value verinin adı float64 de tipi
	if value>=rule.Min && value<=rule.Max{
		return "pass"
	} else {
		return "fail"
	}
}


func mission2(){
	rules := []Rule{
		{Feature: "temperature_mean", Min: 5.0, Max: 25.0,Critical: true},
		{Feature: "precipitation", Min: 300.0,Max:800.0,Critical: false},
	}
	values := []float64{17.4, 250.0}
	for i:=0;i<len(rules);i++{
		checkRule2(rules[i], values[i])
	}
}

func checkRule2(rules Rule,value float64) {
	if value >= rules.Min && value <= rules.Max {
		fmt.Printf("%s -> pass, critical: %t\n",rules.Feature,rules.Critical)				// Not: checkRule2 aynı karşılaştırmayı tekrar yazıyor.	 // Daha temiz yol: içinde checkRule(rule, value) çağırıp sonucu yazdırmak.
	} else {
		fmt.Printf("%s -> fail, critical: %t\n",rules.Feature,rules.Critical)
	}
}