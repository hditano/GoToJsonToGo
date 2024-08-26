package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Product struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Category string  `json:"category"`
}

func main() {
	example := `[
    {
        "name": "Laptop",
        "price": 999.99,
        "category": "Electronics"
    },
    {
        "name": "Coffee Maker",
        "price": 49.99,
        "category": "Home Appliances"
    },
    {
        "name": "Desk Chair",
        "price": 89.99,
        "category": "Furniture"
    },
    {
        "name": "Wireless Mouse",
        "price": 25.50,
        "category": "Electronics"
    }
]`

	p := &Product{"Thinkpad T480", 900, "laptop"}

	// Marshalling JSON - Golang to JSON-
	jsonMarshall, error := json.Marshal(p)
	if error != nil {
		log.Fatal(error)
	}
	// check if Json is valid
	validateMashall := json.Valid([]byte(jsonMarshall))

	fmt.Printf("Is Valid: %v\n", validateMashall)

	fmt.Printf("el Json es %v\n", string(jsonMarshall))

	// Slice of Product to populate with Json-
	var data []Product

	validate := json.Valid([]byte(example))
	fmt.Printf("Is Valid: %v\n", validate)

	// Unmarshalling JSON file - JSON to Golang
	err := json.Unmarshal([]byte(example), &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", data)
}
