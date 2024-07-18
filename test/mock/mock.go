package mock

import (
	"encoding/json"
	"fmt"
	"go_tests/internal/entities"
	"io"
	"os"
)

type CalculatorMock struct {
	entities.Calculator
	Operation string `json:"operation"`
}

type Mock struct {
	Values []CalculatorMock
}

type IMock interface {
	Load()
}

func (m *Mock) Load() {
	path := `C:\Users\Wesley Martins\Projects\go_test\test\mock\singleton.json`

	file, err := os.Open(path)

	if err != nil {
		fmt.Println("ERROR TO LOAD MOCK")
		panic(err)
	}

	bytes, _ := io.ReadAll(file)

	entity := &[]CalculatorMock{}

	json.Unmarshal(bytes, entity)

	m.Values = *entity

	fmt.Println("MOCK LOAD SUCCESS!")
}
