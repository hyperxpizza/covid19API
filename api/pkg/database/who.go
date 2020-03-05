package database

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (db *Database) GetAllWHO(w http.ResponseWriter, r *http.Request) {

	result, err := db.Query("SELECT * FROM csse_recovered")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	SendDataWHO(result, w)
}

func (db *Database) GetAllWHOByCountry(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	result, err := db.Query("SELECT * FROM who WHERE Country=?", params["country"])
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	SendDataWHO(result, w)
}

func (db *Database) GetALLWHOByState(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	result, err := db.Query("SELELCT * FROM who WHERE State=?", params["state"])
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	SendDataWHO(result, w)
}

func (db *Database) GetTotalWHOByCountry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var data int

	result := db.QueryRow("SELECT Feb242020 FROM who WHERE Country=?", params["country"])
	err := result.Scan(&data)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode(data)
}
