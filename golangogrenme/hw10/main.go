package main

import (
    "encoding/json"
    "fmt"
    "os"
	"errors"
)

type CropDefinition struct{
	Crop 	string `json:"crop"`
	Version	string `json:"version"`
	Rules []Rule `json:"rules"`

}

type Rule struct{
	Feature string `json:"feature"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	Critical bool `json:"critical"`
}


func main(){
	data, err := os.ReadFile("wheat.json")
	if err != nil{
		fmt.Println("HATA: Dosya okunamadı ",err)
		return
	}


	var crop CropDefinition // oluşturduğumuz cropdefinition a json verilerini atacağız
	err = json.Unmarshal(data, &crop)

	if err != nil{
		fmt.Println("HATA: json çözümlenemedi ",err)
		return
	}

	fmt.Println("Ürün:",crop.Crop)
	fmt.Println("Versiyon:",crop.Version)				// özellikleri yazdırdık
	fmt.Println("Kural sayisi",len(crop.Rules))

	for _,rule := range crop.Rules{
		err := validateRule(rule)
		if err != nil{
			fmt.Println("HATA: ",err)
		} else{
			fmt.Println(rule.Feature,":geçerli")
			fmt.Println(rule.Feature, "kritik mi:", rule.Critical)
		}
	}
}



func validateRule(rule Rule) error {
	if rule.Feature == "" {
		return errors.New("feature adi bos")
	}																																				//validaterule un daha kısa yazılışı
	if rule.Min > rule.Max {
		return fmt.Errorf("kural %s gecersiz: min (%.1f) > max (%.1f)", rule.Feature, rule.Min, rule.Max)
	}
	return nil
}

