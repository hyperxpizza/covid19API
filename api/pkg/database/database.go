package database

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Database struct {
	*sql.DB
}

func SendDataCSSE(result *sql.Rows, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	var dataTable []DataCSSE

	for result.Next() {
		var data DataCSSE
		err := result.Scan(&data.ID, &data.State, &data.Country, &data.Lat, &data.Long, &data.Jan22_2020, &data.Jan23_2020, &data.Jan24_2020, &data.Jan25_2020, &data.Jan26_2020, &data.Jan27_2020, &data.Jan28_2020, &data.Jan29_2020, &data.Jan30_2020, &data.Jan31_2020, &data.Feb01_2020, &data.Feb02_2020, &data.Feb03_2020, &data.Feb04_2020, &data.Feb05_2020, &data.Feb06_2020, &data.Feb07_2020, &data.Feb08_2020, &data.Feb09_2020, &data.Feb10_2020, &data.Feb11_2020, &data.Feb12_2020, &data.Feb13_2020, &data.Feb14_2020, &data.Feb15_2020, &data.Feb16_2020, &data.Feb17_2020, &data.Feb18_2020, &data.Feb19_2020, &data.Feb20_2020, &data.Feb21_2020, &data.Feb22_2020, &data.Feb23_2020, &data.Feb24_2020, &data.Feb25_2020, &data.Feb26_2020, &data.Feb27_2020, &data.Feb28_2020, &data.Feb29_2020, &data.Mar01_2020, &data.Mar02_2020, &data.Mar03_2020)
		if err != nil {
			panic(err.Error())
		}

		dataTable = append(dataTable, data)
	}

	json.NewEncoder(w).Encode(dataTable)
}

func SendDataWHO(result *sql.Rows, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	var dataTable []DataWHO
	for result.Next() {
		var data DataWHO
		err := result.Scan(&data.ID, &data.State, &data.Country, &data.Jan21_2020, &data.Jan22_2020, &data.Jan23_2020, &data.Jan24_2020, &data.Jan25_2020, &data.Jan26_2020, &data.Jan27_2020, &data.Jan28_2020, &data.Jan29_2020, &data.Jan30_2020, &data.Jan31_2020, &data.Feb01_2020, &data.Feb02_2020, &data.Feb03_2020, &data.Feb04_2020, &data.Feb05_2020, &data.Feb06_2020, &data.Feb07_2020, &data.Feb08_2020, &data.Feb09_2020, &data.Feb10_2020, &data.Feb11_2020, &data.Feb12_2020, &data.Feb13_2020, &data.Feb14_2020, &data.Feb15_2020, &data.Feb16_2020, &data.Feb17_2020, &data.Feb18_2020, &data.Feb19_2020, &data.Feb20_2020, &data.Feb21_2020, &data.Feb22_2020, &data.Feb23_2020, &data.Feb24_2020)
		if err != nil {
			panic(err.Error())
		}

		dataTable = append(dataTable, data)
	}

	json.NewEncoder(w).Encode(dataTable)
}
