package main

import (
	"log"
	handler "mock-json-generator/api"
	"net/http"
)

func main() {

	/* jsonData := []byte(`[
			{
					"name": "name",
					"type": "string",
					"tag": "{firstname}"
			},{
					"name": "create_at",
					"type": "date"
			},
			{
					"name": "age",
					"type": "number",
					"tag": "{number:18,65}"
			},
			{
					"name": "address",
					"type": "object",
					"fields": [
							{
									"name": "street",
									"type": "string",
									"tag": "{street}"
							},
							{
									"name": "city",
									"type": "string",
									"tag": "{city}"
							}
					]
			}
	]`)

	fakeData, err := pkg.ProcessJsonString(jsonData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fakeData)
	*/
	http.HandleFunc("POST /", handler.Handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
