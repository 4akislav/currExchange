package main

import (
	"encoding/json"
	"fmt"
	"genesis/internal"
	"genesis/internal/models"
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fmt.Println("Server started on :8090")

	mux.HandleFunc("/usd/", func(rw http.ResponseWriter, r *http.Request) {

		resp, err := internal.GetRate(rw)
		if err != nil {
			http.Error(rw, "Failed to fetch data from API", http.StatusInternalServerError)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(rw, "Failed to read API response body", http.StatusInternalServerError)
			return
		}

		var data []models.Currency

		if err := json.Unmarshal(body, &data); err != nil {
			fmt.Println("Can not unmarshal JSON", err)
		}

		fmt.Fprintf(rw, "Купівля USD: %.2f\nПродаж USD: %.2f", data[1].Buy, data[1].Sale)

	})

	http.ListenAndServe("localhost:8090", mux)
}
