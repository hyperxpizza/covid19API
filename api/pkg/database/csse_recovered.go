package database

import (
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
