package main

import (
	"math"
	"testing"
)

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-6
}

func TestSmSliceInt(t *testing.T) {
	nums := []int{1, 2, 3}
	expected := 6
	result := sumSliceInt(nums)
	if result != expected {
		t.Errorf("Ошибка! Результат = %d; Должно быть = %d", result, expected)
	}
}

func TestFindFrequencyNum(t *testing.T) {
	nums := []int{1, 2, 2, 3, 3}
	expected := 2
	result := findFrequencyNum(nums)
	if result != expected {
		t.Errorf("Ошибка! Результат = %d; Должно быть = %d", result, expected)
	}
}

func TestCalculateMean(t *testing.T) {
	nums := []int{1, 2, 2, 3, 3}
	expected := 2.2
	result := calculateMean(nums)
	if !almostEqual(result, expected) {
		t.Errorf("Ошибка! Результат = %.7f; Должно быть = %.7f", result, expected)
	}
}

func TestSquaredDeviationsSum(t *testing.T) {
	nums := []int{1, 2, 2, 3, 3}
	expected := 2.8
	result := squaredDeviationsSum(nums, calculateMean(nums))
	if !almostEqual(result, expected) {
		t.Errorf("Ошибка! Результат = %.7f; Должно быть = %.7f", result, expected)
	}
}

func TestPopulationSD(t *testing.T) {
	nums := []int{1, 2, 2, 3, 3}
	expected := 0.748331
	result := populationSD(nums, calculateMean(nums))
	if !almostEqual(result, expected) {
		t.Errorf("Ошибка! Результат = %.7f; Должно быть = %.7f", result, expected)
	}
}

func TestSampleSD(t *testing.T) {
	nums := []int{1, 2, 2, 3, 3}
	expected := 0.83666
	result := sampleSD(nums, calculateMean(nums))
	if !almostEqual(result, expected) {
		t.Errorf("Ошибка! Результат = %.7f; Должно быть = %.7f", result, expected)
	}
}
