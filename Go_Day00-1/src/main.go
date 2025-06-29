package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Функция читает из стандартного потока
func readInput() string {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	return input
}

// Разбивает строку на лексемы и проверяет что каждый элемент инт.
// Сортирует по возрастанию.
func checkInput(in string) ([]int, error) {
	var numbers []int
	sliceStr := strings.Fields(in)
	for _, str := range sliceStr {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}
	sort.Ints(numbers)
	return numbers, nil
}

// Суммирование массива чисел
func sumSliceInt(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// Функция нахождения частотного числа
func findFrequencyNum(nums []int) int {
	var count int
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
		if freq[num] > count {
			count = freq[num]
		}
	}
	mostCommon := make([]int, 0, 5)
	for k, v := range freq {
		if v == count {
			mostCommon = append(mostCommon, k)
		}
	}
	sort.Ints(mostCommon)
	return mostCommon[0]
}

// Считает средние арифметическое
func calculateMean(nums []int) float64 {
	return float64(sumSliceInt(nums)) / float64(len(nums))
}

// Сумма квадратичного отклонения всех элементов
func squaredDeviationsSum(nums []int, mean float64) float64 {
	var output float64
	for _, v := range nums {
		output += math.Pow(float64(v)-mean, 2)
	}
	return output
}

// Стандартное отклонение по генеральной совокупности
func populationSD(nums []int, mean float64) float64 {
	return math.Sqrt(squaredDeviationsSum(nums, mean) / float64(len(nums)))
}

// Стандартное отклонение по выборке
func sampleSD(nums []int, mean float64) float64 {
	return math.Sqrt(squaredDeviationsSum(nums, mean) / float64(len(nums)-1))
}

func main() {
	var mean, median, sd float64
	var mode int
	input := readInput()
	numbers, err := checkInput(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	// Получаем среднее значение
	mean = calculateMean(numbers)
	// Находим медиану слайса.
	len := len(numbers)
	// Если длина слайса нечетная то len(slice)/2+1
	if len%2 > 0 {
		median = float64(numbers[int(len/2)+1])
	} else {
		// иначе len(slice)/2 + len(slice)/2-1
		middle := len / 2
		median = float64(numbers[middle-1]+numbers[middle]) / 2
	}
	// Вычисление частотного числа в слайсе
	mode = findFrequencyNum(numbers)
	// Стандартное отклонение. Буду использовать Стандартное отклонение по генеральной совокупности.
	sd = populationSD(numbers, mean)
	// Просто вывод. Временно
	fmt.Printf(
		"Mean: %.1f\nMedian: %.1f\nMode: %d\nSD: %.2f\n",
		mean, median, mode, sd,
	)
}
