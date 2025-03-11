package main

import (
	"log"
	"mock-json-generator/pkg"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	jsonData := []byte(`[
		{
				"name": "name",
				"type": "string"
		}
	]`)
	result, err := pkg.ProcessJsonString(jsonData)
	if err != nil {
		log.Fatal(err)
	}
	if !strings.Contains(result, "name") {
		log.Fatal("test failed")
	}
	m.Run()
}
