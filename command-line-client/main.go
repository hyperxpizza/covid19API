package main

import (
	"encoding/json"
	"flag"
	"net/http"
	"time"

	"github.com/hyperxpizza/models"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}

var dataset string
var table string

func setFlags() {
	flag.StringVar(&dataset, "ds", "who", "Choose which dataset you want to use, CSSE / WHO")
	flag.StringVar(&table, "t", "", "!ONLY FOR CSSE! Specify the table name: confirmed/deaths/recovered")

	flag.Parse()
}

func getJson(url string, target interface{}) error {
	r, err := client.Get((url))
	if err != nil {
		return err
	}

	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func main() {
	setFlags()

	if dataset == "csse" || dataset == "Csse" || dataset == "CSSE" {
		data := models.DataCSSE{}
		if table == "confirmed" {
			url := "http://localhost:8888/csse/confirmed"
			err := getJson(url, &data)
			if err != nil {
				panic(err.Error())
			}
		}
	}
}
