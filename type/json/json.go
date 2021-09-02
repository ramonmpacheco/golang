package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	//struct to json
	p1 := product{
		ID:    1,
		Name:  "Tv",
		Price: 1000.0,
		Tags:  []string{"Good", "Light"},
	}

	p1json, _ := json.Marshal(p1)

	fmt.Println(string(p1json))

	// struct to json
	var p2 product
	jsonString := string(p1json)
	json.Unmarshal([]byte(jsonString), &p2)

	fmt.Println(p2.Tags[1])
}
