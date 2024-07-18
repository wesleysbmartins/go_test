package automated

import (
	"encoding/json"
	"fmt"
	"go_tests/internal/usecases"
	"go_tests/test/mock"
	result_test "go_tests/test/result"
	"testing"
)

func TestCalculatorUseCase(t *testing.T) {
	mocks := &mock.Mock{}
	mocks.Load()

	fmt.Println("START AUTOMATED TEST")

	for _, mock := range mocks.Values {
		dto := &usecases.CalculatorDTO{
			A:         mock.A,
			B:         mock.B,
			Operation: mock.Operation,
		}

		result, err := usecases.Calculator(*dto)

		jsonMock, _ := json.Marshal(mock)
		jsonResult, _ := json.Marshal(result)

		if err != nil {
			msg := fmt.Sprintf("\nERROR\nEXPECTED: %s\nRESULT: %s\nERROR: %s", jsonMock, jsonResult, err.Error())
			result_test.Error(msg)
			t.Error(msg)
		} else {
			msg := fmt.Sprintf("\nSUCCESS\nEXPECTED: %s\nRESULT: %s", jsonMock, jsonResult)
			result_test.Success(msg)
		}
	}

	fmt.Println("FINISH AUTOMATED TEST")
}
