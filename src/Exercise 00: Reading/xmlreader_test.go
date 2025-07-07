package main

import (
	"os"
	"testing"
)

func TestXMLReader_Read(t *testing.T) {
	xmlData := `<recipes>
        <cake>
            <name>Test Cake</name>
            <stovetime>10 min</stovetime>
            <ingredients>
                <item>
                    <itemname>Flour</itemname>
                    <itemcount>2</itemcount>
                    <itemunit>cups</itemunit>
                </item>
            </ingredients>
        </cake>
    </recipes>`

	tmpFile, err := os.CreateTemp("", "test*.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Write([]byte(xmlData))
	tmpFile.Close()

	reader := XMLReader{}
	data, err := reader.Read(tmpFile.Name())
	if err != nil {
		t.Errorf("Ошибка: %v", err)
	}

	if len(data.Cakes) != 1 {
		t.Errorf("1 пирог != %d", len(data.Cakes))
	}

	cake := data.Cakes[0]
	if cake.Name != "Test Cake" {
		t.Errorf("Ожидался 'Test Cake', получен '%s'", cake.Name)
	}
}

func TestXMLReader_InvalidXML(t *testing.T) {
	invalid := `<recipes><cake><name>Cake</name><stovetime></cake></recipes>` // некорректный XML

	tmpFile, err := os.CreateTemp("", "bad*.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Write([]byte(invalid))
	tmpFile.Close()

	reader := XMLReader{}
	_, err = reader.Read(tmpFile.Name())
	if err == nil {
		t.Error("некорректный XML")
	}
}

func TestXMLReader_EmptyFile(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "empty*.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	reader := XMLReader{}
	_, err = reader.Read(tmpFile.Name())
	if err == nil {
		t.Error("некорректный XML")
	}
}
