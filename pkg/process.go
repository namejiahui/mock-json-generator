package pkg

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

type Field struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Tag    string  `json:"tag,omifieldsty"`
	Fields []Field `json:"fields,omitempty"`
}

func ProcessJsonString(jsonData []byte) (string, error) {
	var fields []Field
	err := json.Unmarshal(jsonData, &fields)
	if err != nil {
		return "", err
	}
	structType := reflect.StructOf(createStructFields(fields))
	structValue := reflect.New(structType).Elem()
	gofakeit.Struct(structValue.Addr().Interface())
	personJSON, err := json.MarshalIndent(structValue.Interface(), "", "  ")
	if err != nil {
		return "", err
	}
	return string(personJSON), nil
}

func createStructFields(fields []Field) (res []reflect.StructField) {
	for _, field := range fields {
		var fieldType reflect.Type
		switch field.Type {
		case "string":
			fieldType = reflect.TypeOf("")
		case "number":
			fieldType = reflect.TypeOf(0.0)
		case "bool":
			fieldType = reflect.TypeOf(false)
		case "date":
			fieldType = reflect.TypeOf(time.Time{})
		case "object":
			fieldType = reflect.StructOf(createStructFields(field.Fields))
		default:
			fieldType = reflect.TypeOf("")
		}
		res = append(res, reflect.StructField{
			Name: strings.ToUpper(string(field.Name[0])) + field.Name[1:],
			Type: fieldType,
			Tag:  reflect.StructTag(`json:"` + field.Name + `" fake:"` + field.Tag + `"`),
		})
	}
	return
}
