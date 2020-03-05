package database

import (
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
