package main

import (
	"fmt"
)

func main() {
	// Initialize the map with a capacity of 10
	// models := make(map[string][]string, 10)

	// // Add models for Ford
	// models["Ford"] = append(models["Ford"], "Mustang", "F-150", "Thunderbolt")

	// // Add models for Chevrolet
	// models["Chevrolet"] = append(models["Chevrolet"], "Camaro", "Silverado", "Corvette")

	// // Add more models for Ford
	// models["Ford"] = append(models["Ford"], "Explorer", "Escape")

	// // Print the map
	// fmt.Println(models)

	map2 := map[string]string{ //creating a map simple
		"99": "milions",
		"56": "sixty",
		"89": "Thaiosand",
		"70": "Hundreds",
		"67": "sixty",
	}

	for id, years := range map2 {
		fmt.Println("The Number is", id, "Value is", years)
	}

	fmt.Println(map2)

	// Define a map with key type int and value type string
	myMap := make(map[string]string) //making a map using make function

	// Assign a value to the key 20
	myMap["20"] = "ABuu"
	myMap["56"] = "Alan"

	// Print the map
	fmt.Println(myMap)

	map_1 := make(map[int]int)

	map_1[20] = 20

	fmt.Println(map_1)
}
