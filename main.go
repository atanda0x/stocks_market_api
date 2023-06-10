package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/atanda0x/stocks_market_api/router"
)

func main() {
	r := router.Router()
	fmt.Println("starting server on the port 9000")

	log.Fatal(http.ListenAndServe(":9000", r))
}
