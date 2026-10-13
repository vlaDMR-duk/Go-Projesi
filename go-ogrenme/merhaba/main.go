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
	fordongu()
	slice()
}
						// buradaki int verinin dönüş tipi, eğer hiçbir şey yazılmazsa C deki gibi void fonksiyonu olur
func topla(a int,b int) int { 		//(a, b int) şeklinde kısa yazım da olur
	return a+b
}

func selamla(isim string) {
	fmt.Printf("Merhaba, %s\n", isim)
}

func fordongu(){
	for i:=0;i<5;i++{
		fmt.Printf("%d\n",i)
	}
}

func slice(){	// c deki dizinin go daki hali
	degerler := []float64{5.0,25.0,4.9,30.0} 		//c deki hali : double degerler[4] = {5.0,25.0,4.9,30.0}
	ilk := degerler[0] 
	fmt.Printf("%.1f \n",ilk)
	degerler[1] = 20
	fmt.Printf("%.1f \n",degerler[1])
	degerler = append(degerler, 12.0)	//append ile ekleme
	fmt.Printf("%.1f \n",degerler[4])
	fmt.Printf("%d \n",len(degerler)) 	//dizi uzunluğu

	for i := 0; i<len(degerler) ; i++ {
		fmt.Printf("%.1f \t",degerler[i])
	}
	fmt.Printf("\n")

	// dizi de geçmek için asıl kullanılan "range"

	for i,deger := range degerler{
		fmt.Printf("i=%d , output=%.2f \t", i, deger)
	}
}