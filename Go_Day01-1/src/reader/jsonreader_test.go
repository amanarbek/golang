package reader

import (
	"os"
	"testing"
)

func TestJSONReader_Read(t *testing.T) {
	jsonData := `{
        "cake": [
            {
                "name": "Test Cake",
                "time": "10 min",
                "ingredients": [
                    {
                        "ingredient_name": "Flour",
                        "ingredient_count": "2",
                        "ingredient_unit": "cups"
                    }
                ]
            }
        ]
    }`

	tmpFile, err := os.CreateTemp("", "test*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Write([]byte(jsonData))
	tmpFile.Close()

	reader := JSONReader{}
	data, err := reader.Read(tmpFile.Name())
	if err != nil {
		t.Errorf("Получена ошибка: %v", err)
	}

	if len(data.Cakes) != 1 {
		t.Errorf("1 пирог != %d", len(data.Cakes))
	}

	cake := data.Cakes[0]
	if cake.Name != "Test Cake" {
		t.Errorf("Ожидался 'Test Cake', получен '%s'", cake.Name)
	}
}

func TestJSONReader_InvalidJSON(t *testing.T) {
	invalid := `{"cake": [ { "name": "Cake" "time": "10 min" }]}` // пропущена запятая

	tmpFile, err := os.CreateTemp("", "bad*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Write([]byte(invalid))
	tmpFile.Close()

	reader := JSONReader{}
	_, err = reader.Read(tmpFile.Name())
	if err == nil {
		t.Error("некорректный JSON")
	}
}

func TestJSONReader_EmptyFile(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "empty*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	reader := JSONReader{}
	_, err = reader.Read(tmpFile.Name())
	if err == nil {
		t.Error("некорректный JSON")
	}
}
