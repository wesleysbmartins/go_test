package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_tests/internal/controllers"
	"go_tests/internal/entities"
	"go_tests/test/mock"
	result_test "go_tests/test/result"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculatorController(t *testing.T) {
	mocks := &mock.Mock{}
	mocks.Load()

	for _, mock := range mocks.Values {
		jsonBody, _ := json.Marshal(mock)
		bodyBuffer := bytes.NewBuffer(jsonBody)

		req, err := http.NewRequest("POST", "http://localhost:8003/calculator", bodyBuffer)
		if err != nil {
			t.Fatal(err)
		}

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(controllers.CalculatorController)
		handler.ServeHTTP(resp, req)

		response := &entities.Calculator{}
		json.NewDecoder(resp.Body).Decode(&response)
		jsonResponse, _ := json.Marshal(response)

		status := resp.Result().StatusCode

		if response.Result != mock.Result || status != http.StatusOK {
			msgErr := fmt.Sprintf("\nERROR\nEXPECTED: %s\nRESPONSE: %s\nERROR: %s\n", jsonBody, jsonResponse, resp.Result().Status)
			result_test.Error(msgErr)
			t.Error(msgErr)
		} else {
			msg := fmt.Sprintf("\nSUCCESS\nMOCKJSON: %s\nRESPONSE: %s", jsonBody, jsonResponse)
			result_test.Success(msg)
		}
	}
}
