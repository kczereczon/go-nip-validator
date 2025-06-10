package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/kczereczon/go-nip-validator/validator"
)

func validateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req map[string]string

	err := json.NewDecoder(r.Body).Decode(&req)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req["type"] == "nip" {
		if !validator.ValidateNip(req["value"]) {
			w.Write([]byte(`{"valid": false}`))
			return
		}

		//json response
		w.Write([]byte(`{"valid": true}`))
		return
	}

	if req["type"] == "pesel" {
		if !validator.ValidatePesel(req["value"]) {
			w.Write([]byte(`{"valid": false}`))
			return
		}

		w.Write([]byte(`{"valid": true}`))
		return
	}

	fmt.Fprintf(w, "Got: %v", req)
}

func main() {
	http.HandleFunc("/validate", validateHandler)

	port := os.Getenv("VALIDATE_PORT")

	if port == "" {
		port = "8080"
	}

	fmt.Fprintf(os.Stdout, "Starting server on port %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
