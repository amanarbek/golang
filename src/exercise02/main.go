package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	oldFile := flag.String("old", "", "Path to the old snapshot file")
	newFile := flag.String("new", "", "Path to the new snapshot file")
	flag.Parse()

	if *oldFile == "" || *newFile == "" {
		fmt.Println("Usage: ./compareFS --old snapshot1.txt --new snapshot2.txt")
		os.Exit(1)
	}

	oldPaths := make(map[string]bool)
	//Читаем старый файл и кладем пути в map
	f1, err := os.Open(*oldFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при открытии старого файла: %v\n", err)
		os.Exit(1)
	}
	defer f1.Close()

	scanner1 := bufio.NewScanner(f1)
	for scanner1.Scan() {
		line := scanner1.Text()
		oldPaths[line] = false
	}
	if err := scanner1.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при чтении старого файла: %v\n", err)
		os.Exit(1)
	}

	// Читаем новый файл и сравниваем
	f2, err := os.Open(*newFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при открытии нового файла: %v\n", err)
		os.Exit(1)
	}
	defer f2.Close()

	scanner2 := bufio.NewScanner(f2)
	for scanner2.Scan() {
		line := scanner2.Text()
		if _, exists := oldPaths[line]; exists {
			oldPaths[line] = true // отмечаем найденный путь true
		} else {
			fmt.Printf("ADDED %s\n", line)
		}
	}
	if err := scanner2.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при чтении нового файла: %v\n", err)
		os.Exit(1)
	}

	// Вывод удаленных путей
	for path, found := range oldPaths {
		if !found {
			fmt.Printf("REMOVED %s\n", path)
		}
	}
}
