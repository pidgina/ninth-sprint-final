package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {

	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"проверяем: 1", 1, 1},
		{"проверяем: 100", 100, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.input)
			assert.Len(t, result, tt.expected)
		})
	}

	assert.Nil(t, generateRandomElements(0))
	assert.Nil(t, generateRandomElements(-1))
	assert.NotNil(t, generateRandomElements(1))

	for _, v := range generateRandomElements(100) {
		assert.Greater(t, v, 0)
	}

}

func TestMaximum(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"слайс с отрицательными числами макимум: -1", []int{-100, -50, -2, -1}, -1},
		{"корректный слайс с одним эллементом, максимум: 99", []int{99}, 99}, //корректно, т.к. мы берем за нулевой к тесту его,
		//а так, как он не больше самого себя цикл просто завершается
		{"корректный слайс, максимум: 100", []int{1, 2, 3, 4, 5, 100, 9}, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.input)
			assert.Equal(t, tt.expected, result)

		})

	}

	assert.Equal(t, 0, maximum(nil))
	assert.Equal(t, 0, maximum([]int{}))

}
