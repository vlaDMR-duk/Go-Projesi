package main

import "fmt"

func main(){
	fmt.Println("hello")
	features := map[string]float64{		//anahtarları string ve değerleri float64 olan bir map
		"temperature_mean": 17.4,
		"precipitation": 250.0,
	}

	fmt.Println(features["temperature_mean"]) 	// çıktı verir
	features["soil_ph"] = 6.8					// yeni anahatar ekler
	features["temperature_mean"] = 18.0			// var olanı değiştirir

	delete(features,"soil_ph")					// değer silme
	fmt.Println(len(features))					// uzunluk verir


	deger := features["soil_ph"]   // soil_ph yok		olmayan anahtarlar için go sıfır değerini verir, python da keyerror alınıyor.
	fmt.Println(deger)              // 0


	deger, ok := features["soil_ph"] 		//comma ok kalıbında ok deger in olup olmamasına bağlı olarak true veya false değerini alır
	if !ok {
		fmt.Println("soil_ph: veri yok -> unknown")
	} else {
		fmt.Println("soil_ph",deger)
	}
	for anahtar,deger := range features{
		fmt.Println(anahtar,deger)
	}
}
