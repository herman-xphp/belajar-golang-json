package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"P001", "name":"Lenovo ThinkBook 14", "price":12000000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]any{
		"id":    "P0001",
		"name":  "Lenovo ThinkBook 14",
		"price": 12000000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
