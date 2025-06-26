package main

import (
	"bufio"
	"fmt"
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

func main() {
	var mean, median, mode, sd float64
	input := readInput()
	numbers, err := checkInput(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	// Получаем среднее значение
	mean = float64(sumSliceInt(numbers)) / float64(len(numbers))
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
}
