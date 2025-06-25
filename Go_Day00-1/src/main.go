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

func main() {
	input := readInput()
	numbers, err := checkInput(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(numbers)
}
