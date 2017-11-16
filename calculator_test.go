package calculator

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	testCases := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 2, 3},
		{-1, 2, 1},
		{1, -2, -1},
		{10, 200, 210},
	}

	for _, tc := range testCases {
		result := Sum(tc.a, tc.b)
		if result != tc.expected {
			t.Error(fmt.Sprintf("Expected %d - Actual: %d", tc.expected, result))
		}
	}
}

func TestSub(t *testing.T) {
	testCases := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 2, -1},
		{-1, 2, -3},
		{1, -2, 3},
		{10, 200, -190},
	}

	for _, tc := range testCases {
		result := Sub(tc.a, tc.b)
		if result != tc.expected {
			t.Error(fmt.Sprintf("Expected %d - Actual: %d", tc.expected, result))
		}
	}
}

func TestMul(t *testing.T) {
	testCases := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 2, 2},
		{-1, 2, -2},
		{1, -4, -4},
		{10, 200, 2000},
	}

	for _, tc := range testCases {
		result := Mul(tc.a, tc.b)
		if result != tc.expected {
			t.Error(fmt.Sprintf("Expected %d - Actual: %d", tc.expected, result))
		}
	}
}

func TestDiv(t *testing.T) {
	testCases := []struct {
		a        int
		b        int
		expected float64
	}{
		{1, 2, 0.5},
		{-1, 2, -0.5},
		{6, -2, -3},
		{10, 200, 0.05},
		{100, 4, 25},
	}

	for _, tc := range testCases {
		result := Div(tc.a, tc.b)
		if result != tc.expected {
			t.Error(fmt.Sprintf("Expected %f - Actual: %f", tc.expected, result))
		}
	}
}