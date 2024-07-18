package unit_test

import (
	"fmt"
	"go_tests/internal/entities"
	"go_tests/test/mock"
	result_test "go_tests/test/result"
	"strconv"
	"testing"
)

func TestCalculatorSum(t *testing.T) {
	mocks := &mock.Mock{}
	mocks.Load()

	sumMock := []mock.CalculatorMock{}

	for _, v := range mocks.Values {
		if v.Operation == "sum" {
			sumMock = append(sumMock, v)
		}
	}

	fmt.Println("Starting SUM Tests")

	for i, mock := range sumMock {
		c := &entities.Calculator{
			A: mock.A,
			B: mock.B,
		}

		c.Sum()

		if c.Result != mock.Result {
			msgErr := fmt.Sprintf("SUM ERROR - EXPECTED: %s + %s = %s RESULT: %s", strconv.Itoa(mock.A), strconv.Itoa(mock.B), strconv.Itoa(mock.Result), strconv.Itoa(c.Result))
			result_test.Error(msgErr)
			t.Error(msgErr)
		} else {
			msg := fmt.Sprintf("MOCK %s OK!", strconv.Itoa(i))
			result_test.Success(msg)
		}
	}
	fmt.Println("SUM TEST FINISHED!")
}

func TestCalculatorSubtraction(t *testing.T) {
	mocks := &mock.Mock{}
	mocks.Load()

	sumMock := []mock.CalculatorMock{}

	for _, v := range mocks.Values {
		if v.Operation == "subtraction" {
			sumMock = append(sumMock, v)
		}
	}

	fmt.Println("Starting SUBTRACTION Tests")

	for i, mock := range sumMock {
		c := &entities.Calculator{
			A: mock.A,
			B: mock.B,
		}

		c.Subtraction()

		if c.Result != mock.Result {
			msgErr := fmt.Sprintf("SUBTRACTION ERROR - EXPECTED: %s - %s = %s RESULT: %s", strconv.Itoa(mock.A), strconv.Itoa(mock.B), strconv.Itoa(mock.Result), strconv.Itoa(c.Result))
			result_test.Error(msgErr)
			t.Error(msgErr)
		} else {
			msg := fmt.Sprintf("MOCK %s OK!", strconv.Itoa(i))
			result_test.Success(msg)
		}
	}
	fmt.Println("SUBTRACTION TEST FINISHED!")
}

func TestCalculatorMultiplication(t *testing.T) {
	mocks := &mock.Mock{}
	mocks.Load()

	sumMock := []mock.CalculatorMock{}

	for _, v := range mocks.Values {
		if v.Operation == "multiplication" {
			sumMock = append(sumMock, v)
		}
	}

	fmt.Println("Starting MULTIPLICATION Tests")

	for i, mock := range sumMock {
		c := &entities.Calculator{
			A: mock.A,
			B: mock.B,
		}

		c.Multiplication()

		if c.Result != mock.Result {
			msgErr := fmt.Sprintf("MULTIPLICATION ERROR - EXPECTED: %s - %s x %s RESULT: %s", strconv.Itoa(mock.A), strconv.Itoa(mock.B), strconv.Itoa(mock.Result), strconv.Itoa(c.Result))
			result_test.Error(msgErr)
			t.Error(msgErr)
		} else {
			msg := fmt.Sprintf("MOCK %s OK!", strconv.Itoa(i))
			result_test.Success(msg)
		}
	}
	fmt.Println("MULTIPLICATION TEST FINISHED!")
}

func TestCalculatorDivisionn(t *testing.T) {
	mocks := &mock.Mock{}
	mocks.Load()

	sumMock := []mock.CalculatorMock{}

	for _, v := range mocks.Values {
		if v.Operation == "division" {
			sumMock = append(sumMock, v)
		}
	}

	fmt.Println("Starting DIVISION Tests")

	for i, mock := range sumMock {
		c := &entities.Calculator{
			A: mock.A,
			B: mock.B,
		}

		c.Division()

		if c.Result != mock.Result {
			msgErr := fmt.Sprintf("DIVISION ERROR - EXPECTED: %s - %s / %s RESULT: %s", strconv.Itoa(mock.A), strconv.Itoa(mock.B), strconv.Itoa(mock.Result), strconv.Itoa(c.Result))
			result_test.Error(msgErr)
			t.Error(msgErr)
		} else {
			msg := fmt.Sprintf("MOCK %s OK!", strconv.Itoa(i))
			result_test.Success(msg)
		}
	}
	fmt.Println("DIVISION TEST FINISHED!")
}
