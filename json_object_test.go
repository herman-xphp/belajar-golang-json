package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Muhammad",
		MiddleName: "Dhabit",
		LastName:   "Masdar",
		Age:        10,
		Married:    false,
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
