package main

import "fmt"

func main() {
	fmt.Println("merhaba go")

	/* Değişkenler*/

	var sicaklik float64 = 17.4
	durum := 17.6

	/*Her iki yazım da aynı anlama geliyor*/

	durum = 18.9 /* = ile := farklı şeyler := ile değişken oluşturulur, = ile değer değiştirilir*/

	//var sayi int        // 0
	//var deger float64   // 0 -> değerin sıfır olması ile değerin olmaması ayırt edilemez
	//var isim string     // "" (boş metin)
	//var kritik bool     // false

	fmt.Printf("Sicaklik: %.1f, durum: %.1f\n", sicaklik, durum)

	if sicaklik >= 5 && sicaklik <= 25 {
		fmt.Printf("pass\n")
	} else { //else in böyle durması zorunlu
		fmt.Printf("fail \n")
	}
}
