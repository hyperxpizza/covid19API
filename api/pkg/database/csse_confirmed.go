package database

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

////////////////////////////// Confirmed
func (db *Database) GetAllCSSEConfirmed(w http.ResponseWriter, r *http.Request) {

	result, err := db.Query("SELECT * FROM csse_confirmed")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	SendDataCSSE(result, w)
}

func (db *Database) GetAllCSSEConfirmedByCountry(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	result, err := db.Query("SELECT * FROM csse_confirmed WHERE Country=?", params["country"])
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	SendDataCSSE(result, w)
}

func (db *Database) GetAllCSSEConfirmedByState(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	result, err := db.Query("SELECT * FROM csse_confirmed WHERE State=?", params["state"])
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	SendDataCSSE(result, w)
}

func (db *Database) GetTotalCSSEConfirmedByCountry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var data int

	result := db.QueryRow("SELECT Mar032020 FROM csse_confirmed WHERE Country=?", params["country"])
	err := result.Scan(&data)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode(data)

}
