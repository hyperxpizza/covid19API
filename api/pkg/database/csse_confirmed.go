package database

import (
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

/*
func (db *Database) GetCSSEConfirmedByCountryAndMonth(w http.ResponseWriter, r *http.Request) {

	result, err := db.Query("select * from information_schema.columns where tablename=csse_confirmed and column_name like 'Jan'")

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	log.Println(result)
}

*/
