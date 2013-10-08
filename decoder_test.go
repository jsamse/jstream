package jstream

import (
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	raw := strings.NewReader(`[{"name":"test","age":30},{"name":"test2","age":40}]`)
	decoder := NewDecoder(raw)
	result := struct {
		Name string
		Age  int
	}{}
	if ok := decoder.Decode(&result); !ok {
		t.Error("Expected successful decode.")
	}
	if err := decoder.Err(); err != nil {
		t.Fatal(err)
	}
	if result.Name != "test" {
		t.Error("Wrong name:", result.Name)
	}
	if result.Age != 30 {
		t.Error("Wrong age:", result.Age)
	}
	if ok := decoder.Decode(&result); !ok {
		t.Error("Expected successful decode.")
	}
	if err := decoder.Err(); err != nil {
		t.Fatal(err)
	}
	if result.Name != "test2" {
		t.Error("Wrong name:", result.Name)
	}
	if result.Age != 40 {
		t.Error("Wrong age:", result.Age)
	}
	if ok := decoder.Decode(&result); ok {
		t.Error("Expected failed decode.")
	}
	if err := decoder.Err(); err != nil {
		t.Fatal(err)
	}
}
