package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ReadJSON() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("requestID readjson", r.Header.Get("request-id"))
		var jsonInput map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&jsonInput)

		if err != nil {
			fmt.Errorf("error  %w", err)
		}

		fmt.Printf("%s", jsonInput["address"])

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(&jsonInput)
		if err != nil {
			fmt.Errorf("error  %w", err)
		}
	}
}

