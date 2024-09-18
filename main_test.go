package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCafe(t *testing.T) {
	type testCase struct {
		name     string
		menu     string
		money    float64
		expected float64
	}

	testCases := []testCase{
		{
			name:     "buy_tea_100_expect_change_50",
			menu:     "tea", // Price 50
			money:    100,
			expected: 50,
		},
		{
			name:     "buy_greentea_100_expect_change_40",
			menu:     "greentea", // Price 60
			money:    100,
			expected: 40,
		},
		{
			name:     "buy_coffee_500_expect_change_430",
			menu:     "coffee", // Price 70
			money:    500,
			expected: 430,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			menu, change := cafe(tc.menu, tc.money)
			assert.Equal(t, tc.menu, menu, fmt.Sprintf("menu %s", menu))
			assert.Equal(t, tc.expected, change, fmt.Sprintf("money %0.0f, expected change %0.0f", tc.money, tc.expected))
		})
	}
}
