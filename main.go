package main

import (
	"encoding/json"
	"fmt"
	"genesis/internal"
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fmt.Println("Server started on :8090")

	mux.HandleFunc("/rate/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("https://api.privatbank.ua/p24api/pubinfo?json&exchange&coursid=5")
		if err != nil {
			http.Error(w, "Failed to fetch data from API", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read API response body", http.StatusInternalServerError)
			return
		}

		var data []internal.Currency

		if err := json.Unmarshal(body, &data); err != nil {
			fmt.Println("Can not unmarshal JSON", err)
		}

		fmt.Fprintf(w, "Купівля USD: %.2f\nПродаж USD: %.2f", data[1].Buy, data[1].Sale)

	})

	http.ListenAndServe("localhost:8090", mux)
}
