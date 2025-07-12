package main

import (
	"flag"
	"log"
	"project/reader"
)

func main() {
	oldFile := flag.String("old", "", "path to original database")
	newFile := flag.String("new", "", "path to stolen database")
	flag.Parse()

	if *oldFile == "" || *newFile == "" {
		log.Fatal("Оба флага --old и --new обязательны")
	}

	oldReader, err := reader.GetReader(*oldFile)
	if err != nil {
		log.Fatal(err)
	}

	newReader, err := reader.GetReader(*newFile)
	if err != nil {
		log.Fatal(err)
	}

	oldDB, err := oldReader.Read(*oldFile)
	if err != nil {
		log.Fatal(err)
	}

	newDB, err := newReader.Read(*newFile)
	if err != nil {
		log.Fatal(err)
	}

	oldDB.CompareRecipes(*newDB)
}
