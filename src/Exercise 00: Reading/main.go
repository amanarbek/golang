package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"path/filepath"
)

func main() {
	file := flag.String("f", "", "Чтобы спарсить файл, укажи до него путь")
	flag.Parse()
	if *file == "" {
		fmt.Println("Использовать флаг -f ?")
		return
	}

	dbReader, err := GetReader(*file)
	if err != nil {
		fmt.Println("Не удалось прочесть файл")
		return
	}

	recipes, err := dbReader.Read(*file)
	if err != nil {
		fmt.Println("Ошибка чтения", err)
		return
	}

	ext := filepath.Ext(*file)
	switch ext {
	case ".json":
		out, _ := xml.MarshalIndent(recipes, "", "    ")
		fmt.Println(xml.Header + string(out))
	case ".xml":
		out, _ := json.MarshalIndent(recipes, "", "    ")
		fmt.Println(string(out))
	default:
		fmt.Println("Неизвестный тип файла")
	}
}
