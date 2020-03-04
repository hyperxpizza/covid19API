package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hyperxpizza/pkg/database"
	_ "github.com/mattn/go-sqlite3"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to index")
}

func connectToDB() *sql.DB {
	db, err := sql.Open("sqlite3", "../covid19.db")
	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	//connect to DB
	db := connectToDB()
	sqlite := database.Database{db}

	defer db.Close()

	router := mux.NewRouter().StrictSlash((true))
	router.HandleFunc("/", index)

	//API
	//CSSE data
	//Confirmed
	router.HandleFunc("/csse/confirmed", sqlite.GetAllCSSEConfirmed).Methods("GET")
	router.HandleFunc("/csse/confirmed/{country}", sqlite.GetAllCSSEConfirmedByCountry).Methods("GET")

	//Deaths
	router.HandleFunc("/csse/deaths", sqlite.GetAllCSSEDeaths).Methods("GET")
	router.HandleFunc("/csse/deaths/{country}", sqlite.GetAllCSSEDeathsByCountry).Methods("GET")

	//Recovered
	router.HandleFunc("/csse/recovered", sqlite.GetAllCSSERecovered).Methods("GET")
	router.HandleFunc("/csse/recovered/{country}")

	//WHO DATA
	router.HandleFunc("/who", sqlite.GetAllWHO).Methods("GET")

	log.Fatal(http.ListenAndServe(":8888", router))
}
