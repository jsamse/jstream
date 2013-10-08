package jstream

import (
	"bytes"
	"testing"
)

func TestEncode(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	encoder := NewEncoder(buffer)
	err := encoder.Encode(struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "test",
		Age:  30,
	})
	if err != nil {
		t.Fatal(err)
	}
	err = encoder.Encode(struct {
		Unit  string `json:"unit"`
		Value int    `json:"value"`
	}{
		Unit:  "km",
		Value: 50,
	})
	if err != nil {
		t.Fatal(err)
	}
	err = encoder.Close()
	if err != nil {
		t.Fatal(err)
	}
	if buffer.String() != `[{"name":"test","age":30},{"unit":"km","value":50}]` {
		t.Error("Wrong result:", buffer.String())
	}
}
