package models

import (
	"encoding/xml"
	"fmt"
)

type Ingredient struct {
	XMLName xml.Name `json:"-" xml:"item"`
	Name    string   `json:"ingredient_name" xml:"itemname"`
	Count   string   `json:"ingredient_count" xml:"itemcount"`
	Unit    string   `json:"ingredient_unit,omitempty" xml:"itemunit"`
}

type Cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

type Recipes struct {
	XMLName xml.Name `json:"-" xml:"recipes"`
	Cakes   []Cake   `json:"cake" xml:"cake"`
}

func (r Recipes) CompareRecipes(other Recipes) {
	orgRes := make(map[string]Cake)
	commonRecipes := make(map[string]Cake)

	for _, cake := range r.Cakes {
		orgRes[cake.Name] = cake
	}

	// Найти добавленные и удалённые торты
	for _, cake := range other.Cakes {
		if _, ok := orgRes[cake.Name]; !ok {
			fmt.Printf("ADDED cake %q\n", cake.Name)
		} else {
			commonRecipes[cake.Name] = cake
		}
	}

	for _, cake := range orgRes {
		if _, ok := commonRecipes[cake.Name]; !ok {
			fmt.Printf("REMOVED cake %q\n", cake.Name)
		}
	}

	// Найти изменения в общих тортах
	for name, cake := range commonRecipes {
		if cake.Time != orgRes[name].Time {
			fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", name, cake.Time, orgRes[name].Time)
		}

		// И сравнение ингредиентов
		orgRes[name].compareIngredients(cake)
	}
}

func (c Cake) compareIngredients(other Cake) {

}
