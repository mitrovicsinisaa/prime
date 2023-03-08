package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/mitrovicsinisaa/prime/primes"
	"github.com/stretchr/testify/assert"
)

func TestPrimesSuccess(t *testing.T) {
	app := fiber.New()
	app.Post("/prime", primes.CheckPrimes)

	testCase := struct {
		name         string
		requestBody  string
		expectedBody []interface{}
		expectedCode int
	}{
		name:         "Valid input",
		requestBody:  `[2,3,4,5]`,
		expectedBody: []interface{}{true, true, false, true},
		expectedCode: http.StatusOK,
	}

	t.Run(testCase.name, func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/prime", strings.NewReader(testCase.requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, testCase.expectedCode, resp.StatusCode)

		var body interface{}
		err = json.NewDecoder(resp.Body).Decode(&body)
		assert.NoError(t, err)
		assert.Equal(t, testCase.expectedBody, body)
	})
}

func TestPrimesError(t *testing.T) {
	app := fiber.New()
	app.Post("/prime", primes.CheckPrimes)

	testCase := struct {
		name         string
		requestBody  string
		expectedBody map[string]interface{}
		expectedCode int
	}{
		name:         "Invalid input",
		requestBody:  `[2,3,"nan"]`,
		expectedBody: map[string]interface{}{"error": "Element on index 2 is not valid"},
		expectedCode: http.StatusBadRequest,
	}

	t.Run(testCase.name, func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/prime", strings.NewReader(testCase.requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, testCase.expectedCode, resp.StatusCode)

		var body interface{}
		err = json.NewDecoder(resp.Body).Decode(&body)
		assert.NoError(t, err)
		assert.Equal(t, testCase.expectedBody, body)
	})
}
