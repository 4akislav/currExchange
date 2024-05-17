package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	exchange_url := "https://api.privatbank.ua/p24api/pubinfo?json&exchange&coursid=5"

	router := gin.Default()
	router.GET("/usd", func(ctx *gin.Context) {
		getUsd(ctx, exchange_url)
	})
	router.ServeHTTP(w, r)
}

func getUsd(ctx *gin.Context, exchange_url string) {
	ctx.String(http.StatusOK, exchange_url)
}
