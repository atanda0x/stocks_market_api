package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/atanda0x/stocks_market_api/router"
)

func main() {
	r := router.Router()
	fmt.Println("starting serveron the port 5432")

	log.Fatal(http.ListenAndServe(":5432", r))
}
