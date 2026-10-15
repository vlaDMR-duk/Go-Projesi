package main

import (
	"errors"
	"fmt"
)


func main(){
	fmt.Println("hello")
	toplam, carpim := hesapla(10,2)			//bu şekilde fonksiyona 2 paramtre yollayabiliyoruz
	fmt.Println(toplam)
	fmt.Println(carpim)
	sonuc, err := bolme(10, 0)
	if err != nil {							// yani hata varsa
		fmt.Println("HATA:", err)
		return
	}
	fmt.Println("Sonuc:", sonuc)
}

func hesapla(a, b float64) (float64,float64){
	return a+b,a*b
}

func bolme(a float64,b float64) (float64,error) {
	if b==0 {
		return 0, errors.New("sifira bölme")				//sıfıra bölünemeyeceğinden ötürü hata alıyoruz, eğer b = 0 olursa sıfır döndürüp hata mesajı olarak sifira bölme alıyoruz
	}
	return a / b, nil
}
