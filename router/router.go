package router

import (
	"github.com/atanda0x/stocks_market_api/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stock", middleware.GetAllStock).Methods("GET", "OPTION")
	router.HandleFunc("/api/newstock", middleware.CreateStock).Methods("POST", "OPTION")
	router.HandleFunc("/api/stock/{id}", middleware.UpdateStock).Methods("PUT", "OPTION")
	router.HandleFunc("/api/deletestock/{id}", middleware.DeleteStock).Methods("DELETE", "OPTION")

}
