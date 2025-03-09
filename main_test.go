package main

import (
	"log"
	"mock-json-generator/pkg"
	"testing"
)

func TestMain(m *testing.M) {
	jsonData := []byte(`[
		{
				"name": "name",
				"type": "string"
		}
	]`)
	_, err := pkg.ProcessJsonString(jsonData)
	if err != nil {
		log.Fatal(err)
	}
	m.Run()
}
