package router

import (
	"github.com/gorilla/mux"
)

func router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/")
	router.HandleFunc("/")
	router.HandleFunc("/")
	router.HandleFunc("/")
	router.HandleFunc("/")

}
