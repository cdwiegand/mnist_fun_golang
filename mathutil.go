package main

import (
	"math"
	"math/rand"
)

func MatrixMultiply(data [][]float64, columnMultipliers []float64) []float64 {
	if len(columnMultipliers) != len(data[0]) {
		panic("Bad matrix!")
	}

	ret := make([]float64, len(data[0]))

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			ret[i] += data[i][j] * columnMultipliers[j]
		}
	}
	return ret
}

func Sigmoid(h_pre float64) float64 {
	float1 := float64(1)
	return float1 / (float1 + math.Pow(math.E, 0-h_pre))
}

func FillWithRandom(rand *rand.Rand, count int) []float64 {
	matrix := make([]float64, count)
	for i2hl1 := 0; i2hl1 < len(matrix); i2hl1++ {
		matrix[i2hl1] = rand.Float64() - 0.5
	}
	return matrix
}

func NewSingleMatrix(countNeurons int) []float64 {
	return make([]float64, countNeurons)
}
func FillWithValueF(value float64, matrix []float64) {
	for i2hl1 := 0; i2hl1 < len(matrix); i2hl1++ {
		matrix[i2hl1] = value
	}
}
func FillWithValueI(value int, matrix []float64) {
	for i2hl1 := 0; i2hl1 < len(matrix); i2hl1++ {
		matrix[i2hl1] = float64(value)
	}
}
