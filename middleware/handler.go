package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/atanda0x/stocks_market_api/models"
	"github.com/gorilla/mux"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateConnection() *sql.DB {
	err := gotdotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to postgres")
	return db

}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatal("Unable to decode the request body. %v", err)
	}

	insertID := insertStock(stock)

	res := response{
		ID:      insertID,
		Message: "stock created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	stock, err := getStock(int64(id))

	if err != nil {
		log.Fatal("Unable to get stock. %v", err)
	}

	json.NewEncoder(w).Encode(stock)
}

func GetAllStock(w http.ResponseWriter, r *http.Request) {
	stocks, err := getAllStock()

	if err != nil {
		log.Fatalf("Unable to get all the stocks. %v")
	}

	json.NewEncoder(w).Encode(stocks)
}
func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string innto int. %v", err)
	}

	var stock models.Stock

	err = json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	updatedRows := updateStock(int64(id), stock)

	msg := fmt.Sprintf("Stock updated sucessfully. total rows/record affected %v", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string to int. %v", err)
	}

	deleteRows := deleteStock(int64(id))

	msg := fmt.Sprintf("Stocks deleted successfully. Total rows/records %v", deleteRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}
