package main

import (
	"encoding/json"
	"fmt"
	"genesis/internal"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fmt.Println("Server started on :8090")

	mux.HandleFunc("/usd/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("https://api.privatbank.ua/p24api/pubinfo?json&exchange&coursid=5")
		if err != nil {
			http.Error(w, "Failed to fetch data from API", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		var data []internal.Currency

		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Failed to decode API response", http.StatusInternalServerError)
			return
		}

		if len(data) == 0 {
			http.Error(w, "No currency data found", http.StatusInternalServerError)
			return
		}

		usdExchange := data[0]

		fmt.Fprintf(w, "Вартість доллара в закупівлю: %.2f.\nВартість доллара на продаж: %.2f.", usdExchange.Buy, usdExchange.Sale)

	})

	http.ListenAndServe("localhost:8090", mux)
}
