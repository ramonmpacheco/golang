package math

import "testing"

const defaultError = "Expected value %v, but found %v."

func TestAvg(t *testing.T) {
	expectedValue := 7.28

	value := Avg(7.2, 9.9, 6.1, 5.9)

	if value != expectedValue {
		t.Errorf(defaultError, expectedValue, value)
	}
}
// to run all tests via terminal: ```go test ./...```