package main

import (
	"cell-shield-go/shieldinformation"
	"log"
)

func main() {

	shieldInformation, err := shieldinformation.GrabShieldInformation("", "1-m3sfTXGwoqsPrZcyJQnJ2FoClwpl67EyPwUtOmwtxo", "B3")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(*shieldInformation)

}
