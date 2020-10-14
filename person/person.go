package person

import (
	"encoding/json"
	"log"
)

type Person struct {
	Name string
	Age  int
}

func ToJson(structure interface{}) []byte {
	bytes, err := json.Marshal(structure)

	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

func GetPerson() Person {
	return Person{
		Name: "Philipp",
		Age:  25,
	}
}
