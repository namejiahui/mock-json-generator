package handler

import (
	"fmt"
	"io"
	"mock-json-generator/pkg"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	jsonResponse, err := pkg.ProcessJsonString(body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, jsonResponse)
}
