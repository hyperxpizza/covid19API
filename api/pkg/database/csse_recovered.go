package database

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//////////////////////////////////////// RECOVERED

func (db *Database) GetAllCSSERecovered(w http.ResponseWriter, r *http.Request) {

	result, err := db.Query("SELECT * FROM csse_recovered")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	SendDataCSSE(result, w)
}

func (db *Database) GetAllCSSERecoveredByCountry(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	result, err := db.Query("SELECT * FROM csse_recovered WHERE Country=?", params["country"])
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	SendDataCSSE(result, w)
}

func (db *Database) GetAllCSSERecoveredByState(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	result, err := db.Query("SELECT * FROM csse_recovered WHERE State=?", params["state"])
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	SendDataCSSE(result, w)
}

func (db *Database) GetTotalCSSERecoveredByCountry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var data int

	result := db.QueryRow("SELECT Mar032020 FROM csse_recovered WHERE Country=?", params["country"])
	err := result.Scan(&data)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode(data)

}
