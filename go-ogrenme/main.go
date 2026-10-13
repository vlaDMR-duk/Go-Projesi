package main

import "fmt"

func main(){
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

	fmt.Printf("Sicaklik: %.1f, durum: %.1f\n",sicaklik,durum)
	sonuc := topla(3,4)
	fmt.Printf("%d\n",sonuc)
	isim := "Demir"
	selamla(isim)
}
						// buradaki int verinin dönüş tipi, eğer hiçbir şey yazılmazsa C deki gibi void fonksiyonu olur
func topla(a int,b int) int { 		//(a, b int) şeklinde kısa yazım da olur
	return a+b
}

func selamla(isim string) {
	fmt.Printf("Merhaba, %s\n", isim)
}