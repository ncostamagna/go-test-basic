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

func TestSumMatrix(t *testing.T) {
	cases := []struct {
		name string
		input []int
		expected int		
	}{
		{
			name: "Sum [1, 2, 3, 4, 5] = 15",
			input: []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name: "Sum [2,3,1,2] = 8",
			input: []int{2,3,1,2},
			expected: 8,
		},
		{
			name: "Sum [1,1,1,1] = 4",
			input: []int{1,1,1,1},
			expected: 4,
		},
		{
			name: "Sum [] = 0",
			input: []int{},
			expected: 0,
		},
		{
			name: "Sum [-10,9] = -1",
			input: []int{-10,9},
			expected: -1,
		},
	}
	
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := Sum(c.input...)
			assert.Equal(t, c.expected, result)
		})
	}
	
}

func TestCheckEmailMatrix(t *testing.T) {
	cases := []struct {
		name string
		input string
		expected bool
	}{
		{
			name: "check valid test@test.com",
			input: "test@test.com",
			expected: true,
		},
		{
			name: "check valid nahuel@test.ai",
			input: "nahuel@test.ai",
			expected: true,
		},
		{
			name: "check invalid nahuel@test",
			input: "nahuel@test",
			expected: false,
		},
		{
			name: "check invalid nahueltest.com",
			input: "nahueltest.com",
			expected: false,
		},
		{
			name: "check invalid nahuel@.com",
			input: "nahuel@.com",
			expected: false,
		},
		{
			name: "check invalid nahuel@.ai",
			input: "nahuel@.ai",
			expected: false,
		},
		{
			name: "check invalid nahuel@.com.ar",
			input: "nahuel@.com.ar",
			expected: false,
		},
		{
			name: "check invalid nahuel@.com.ar.com",
			input: "nahuel@.com.ar.com",
			expected: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := CheckEmail(c.input)
			assert.Equal(t, c.expected, result)
		})
	}
}