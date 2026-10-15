package main

import (
    "encoding/json"
    "fmt"
    "os"
)
type CropDefinition struct {
    Crop    string `json:"crop"`
    Version string `json:"version"`
    Rules   []Rule `json:"rules"`
}

type Rule struct {
    Feature  string  `json:"feature"`
    Min      float64 `json:"min"`
    Max      float64 `json:"max"`
    Critical bool    `json:"critical"`
}


func main() {

    data, err := os.ReadFile("wheat.json")
    if err != nil {
        fmt.Println("HATA: dosya okunamadi:", err)
        return
    }

    var crop CropDefinition
    err = json.Unmarshal(data, &crop)
    if err != nil {
        fmt.Println("HATA: JSON cozumlenemedi:", err)
        return
    }

    fmt.Println("Urun:", crop.Crop, "Versiyon:", crop.Version)
    fmt.Println("Kural sayisi:", len(crop.Rules))
}