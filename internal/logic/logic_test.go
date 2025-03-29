package logic

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	result := Sum(2, 3)
	expected := 5

	assert.Equal(t, expected, result)

	result = internalSum(1, 2, 3, 4, 5)
	expected = 15

	assert.Equal(t, expected, result)
}

func TestCheckEmail(t *testing.T) {
	result := CheckEmail("test@test.com")
	expected := true

	assert.Equal(t, expected, result)
}