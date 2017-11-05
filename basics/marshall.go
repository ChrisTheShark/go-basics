package basics

import (
	"fmt"
	"encoding/json"
	"log"
)

type Shark struct {
	Length float64 `json:"length"`
	CommonName string `json:"name"`
	ScientificName string `json:"latin"`
	id int
}

func main() {
	marshall()
	unmarshall()
}

func marshall() {
	tiger := &Shark{13.4, "Tiger Shark", "Galeocerdo cuvier", 1}
	fmt.Println(tiger)

	if bs, err := json.Marshal(tiger); err == nil {
		fmt.Println(string(bs))
	}
}

func unmarshall() {
	var white Shark
	jsonString := []byte(`{"length":18.4,"name":"Great White Shark","latin":"Carcharodon carcharias"}`)

	if err := json.Unmarshal(jsonString, &white); err != nil {
		log.Fatal(err)
	}

	fmt.Println(white)
}