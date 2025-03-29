package logic

import (
	"testing"
)

func TestSum(t *testing.T) {
	result := Sum(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Sum(2, 3) = %d; expected %d", result, expected)
	}

	result = internalSum(1, 2, 3, 4, 5)
	expected = 15

	if result != expected {
		t.Errorf("Sum(1, 2, 3, 4, 5) = %d; expected %d", result, expected)
	}
}

func TestCheckEmail(t *testing.T) {
	result := CheckEmail("test@test.com")
	expected := true

	if result != expected {
		t.Errorf("CheckEmail(\"test@test.com\") = %t; expected %t", result, expected)
	}
}