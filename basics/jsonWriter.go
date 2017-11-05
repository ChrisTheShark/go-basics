package basics

import (
	"encoding/json"
	"fmt"
	"strings"
)

type StdoutWriter struct {}

func (writer StdoutWriter) Write(p []byte) (n int, err error) {
	n, _ = fmt.Println(string(p))
	return n, nil
}

type Shape struct {
	Sides int `json:"sides"`
	AngleSum float64 `json:"sum_angles"`
}

func main() {
	shape := Shape{3, 180}
	json.NewEncoder(StdoutWriter{}).Encode(shape)

	var shape2 Shape
	json.NewDecoder(strings.NewReader(`{"sides": 3, "sum_angles": 180}`)).Decode(&shape2)
	fmt.Println(shape2)
}