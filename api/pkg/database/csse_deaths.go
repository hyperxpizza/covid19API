package database

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

/////////////////////////////////////DEATHS

func (db *Database) GetAllCSSEDeaths(w http.ResponseWriter, r *http.Request) {
	result, err := db.Query("SELECT * FROM csse_deaths")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	SendDataCSSE(result, w)
}

func (db *Database) GetAllCSSEDeathsByCountry(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	result, err := db.Query("SELECT * FROM csse_deaths WHERE Country=?", params["country"])
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	SendDataCSSE(result, w)
}

func (db *Database) GetAllCSSEDeathsByState(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	result, err := db.Query("SELECT * FROM csse_deaths WHERE State=?", params["state"])
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	SendDataCSSE(result, w)
}

func (db *Database) GetTotalCSSEDeathsByCountry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var data int

	result := db.QueryRow("SELECT Mar032020 FROM csse_deaths WHERE Country=?", params["country"])
	err := result.Scan(&data)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode(data)

}
