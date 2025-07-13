package models

import (
	"encoding/xml"
	"fmt"
	"strings"
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
	ings := make(map[string]Ingredient)
	commonIngs := make(map[string]Ingredient)

	for _, ing := range c.Ingredients {
		ings[strings.ToLower(ing.Name)] = ing
	}

	// Сравнивает ингредиенты
	for _, ing := range other.Ingredients {
		keyToLow := strings.ToLower(ing.Name)
		if _, ok := ings[keyToLow]; !ok {
			fmt.Printf("ADDED ingredient \"%s\" for cake  \"%s\"\n", ing.Name, other.Name)
		} else {
			// Добавить общий ингредиент
			commonIngs[keyToLow] = ing
		}
		if ings[keyToLow].Count != ing.Count {
			fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n",
				ing.Name, other.Name, ing.Count, ings[keyToLow].Count)
		}
		// Единицы измерения могут быть опущены. Проверка на пустую строку обязательно
		if ings[keyToLow].Unit != ing.Unit {
			if ings[keyToLow].Unit == "" {
				fmt.Printf("ADDED unit \"%s\" for ingredient \"%s\" for cake  \"%s\"\n", ing.Unit, ing.Name, other.Name)
			} else if ing.Unit == "" {
				fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake  \"%s\"\n", ings[keyToLow].Unit, ings[keyToLow].Name, c.Name)
			} else {
				fmt.Printf("CHANGED unit for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n",
					ings[keyToLow].Name, c.Name, ing.Unit, ings[keyToLow].Unit)
			}
		}
	}
	// Доп проход для вывода удаленных ингредиентов
	for _, ing := range ings {
		if _, ok := commonIngs[ing.Name]; !ok {
			fmt.Printf("REMOVED ingredient \"%s\" for cake  \"%s\"\n", ing.Name, other.Name)
		}
	}
}
