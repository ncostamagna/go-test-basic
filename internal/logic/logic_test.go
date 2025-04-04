package logic_test

import (
	"testing"
	"fmt"
	"os"
	"github.com/stretchr/testify/assert"
	"github.com/ncostamagna/go-test-basic/internal/logic"
)

func TestMain(m *testing.M) {
	fmt.Println("Starting tests")
	exitCode := m.Run()
	fmt.Println("Tests finished")
	os.Exit(exitCode)
}

func TestSum(t *testing.T) {
	got := logic.Sum(2, 3)
	expected := 5

	assert.Equal(t, expected, got)

	got = logic.Sum(1, 2, 3, 4, 5)
	expected = 15
	assert.Equal(t, expected, got)
}

func TestCheckEmail(t *testing.T) {
	email := "test@test.com"
	got := logic.CheckEmail(email)
	assert.True(t, got)

	invalidEmail := "testtest.com.ar"
	got = logic.CheckEmail(invalidEmail)
	assert.False(t, got)
}


func TestSumMatrix(t *testing.T) {
	cases := []struct {
		name string
		input []int
		expected int
	}{
		{
			name: "Sum of 1, 2, 3",
			input: []int{1, 2, 3},
			expected: 6,
		},
		{
			name: "Sum of 4, 5, 6",
			input: []int{4, 5, 6},
			expected: 15,
		},
		{
			name: "Sum of 7, 8, 9",
			input: []int{7, 8, 9},
			expected: 24,
		},
		{
			name: "Sum of 10, -11, -12",
			input: []int{10, -11, -12},
			expected: -13,
		},
		{
			name: "Sum of empty",
			input: []int{},
			expected: 0,
		},
	}
	
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := logic.Sum(c.input...)
			assert.Equal(t, c.expected, got)
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
			name: "check email test@test.com",
			input: "test@test.com",
			expected: true,
		},
		{
			name: "check email testtest.com.ar",
			input: "testtest.com.ar",
			expected: false,
		},
		{
			name: "check email test@test",
			input: "test@test",
			expected: false,
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
		{
			name: "check valid nahuel@test.com.ar",
			input: "nahuel@test.com.ar",
			expected: true,
		},
		{
			name: "check empty email",
			input: "",
			expected: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := logic.CheckEmail(c.input)
			assert.Equal(t, c.expected, got)
		})
	}
}

func BenchmarkSum(b *testing.B) {
	for b.Loop() {
		_ = logic.Sum(1, 2, 3, 4, 5)
	}
}

func BenchmarkCheckEmail(b *testing.B) {
	for b.Loop() {
		_ = logic.CheckEmail("")
	}
}

