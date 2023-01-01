package main

import "fmt"

func main() {
	groupCity := map[int][]string{
		10:   {"Деревня", "Село"},
		100:  {"Город", "Большой город"},
		1000: {"Миллионик"},
	}
	cityPopulation := map[string]int{
		"Село":          10,
		"cело":          100,
		"Миллионик":     1000,
		"Город":         100,
		"Большой город": 100,
	}

	for key := range cityPopulation {
		var exist bool
		for _, city := range groupCity[100] {
			if city == key {
				exist = true
				break
			}
		}
		if !exist {
			delete(cityPopulation, key)
		}
	}

	fmt.Println(cityPopulation)
}
