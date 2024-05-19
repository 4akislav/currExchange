package internal

import "net/http"

func GetRate(rw http.ResponseWriter) (*http.Response, error) {

	resp, err := http.Get("https://api.privatbank.ua/p24api/pubinfo?json&exchange&coursid=5")
	if err != nil {
		return nil, err
	}

	return resp, nil
}
