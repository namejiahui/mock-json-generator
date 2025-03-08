package main

import (
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	jsonData := []byte(`[
		{
				"name": "name",
				"type": "string"
		}
	]`)
	_, err := processJsonString(jsonData)
	if err != nil {
		log.Fatal(err)
	}
	m.Run()
}
