package reading

import (
	"testing"
)

func TestGetReader(t *testing.T) {
	tests := []struct {
		filename   string
		wantErr    bool
		wantReader string
	}{
		{"data.json", false, "JSONReader"},
		{"data.xml", false, "XMLReader"},
		{"data.txt", true, ""},
	}

	for _, tt := range tests {
		r, err := GetReader(tt.filename)
		if (err != nil) != tt.wantErr {
			t.Errorf("GetReader(%q) error = %v, wantErr %v", tt.filename, err, tt.wantErr)
			continue
		}

		if err == nil {
			switch r.(type) {
			case JSONReader:
				if tt.wantReader != "JSONReader" {
					t.Errorf("Должен быть JSONReader")
				}
			case XMLReader:
				if tt.wantReader != "XMLReader" {
					t.Errorf("Должен быть XMLReader")
				}
			default:
				t.Errorf("Не знаю такой тип")
			}
		}
	}
}
