package usecases

import (
	"fmt"
	"go_tests/internal/entities"
)

func Calculator(dto CalculatorDTO) (entities.Calculator, error) {
	c := &entities.Calculator{
		A: dto.A,
		B: dto.B,
	}

	var err error

	if dto.Operation == "sum" {
		c.Sum()
	} else if dto.Operation == "subtraction" {
		c.Subtraction()
	} else if dto.Operation == "multiplication" {
		c.Multiplication()
	} else if dto.Operation == "division" {
		c.Division()
	} else {
		err = fmt.Errorf("Unknown Operation")
	}

	return *c, err
}
