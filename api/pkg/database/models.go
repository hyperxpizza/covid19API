package database

import "database/sql"

type DataCSSE struct {
	ID         int            `json:"id"`
	State      sql.NullString `json:"state"`
	Country    sql.NullString `json:"country"`
	Lat        float32        `json:"lat"`
	Long       float32        `json:"long"`
	Jan22_2020 sql.NullString `json:"01/22/2020"`
	Jan23_2020 sql.NullString `json:"01/23/2020"`
	Jan24_2020 sql.NullString `json:"01/24/2020"`
	Jan25_2020 sql.NullString `json:"01/25/2020"`
	Jan26_2020 sql.NullString `json:"01/26/2020"`
	Jan27_2020 sql.NullString `json:"01/27/2020"`
	Jan28_2020 sql.NullString `json:"01/28/2020"`
	Jan29_2020 sql.NullString `json:"01/29/2020"`
	Jan30_2020 sql.NullString `json:"01/30/2020"`
	Jan31_2020 sql.NullString `json:"01/31/2020"`
	Feb01_2020 sql.NullString `json:"02/01/2020"`
	Feb02_2020 sql.NullString `json:"02/02/2020"`
	Feb03_2020 sql.NullString `json:"02/03/2020"`
	Feb04_2020 sql.NullString `json:"02/04/2020"`
	Feb05_2020 sql.NullString `json:"02/05/2020"`
	Feb06_2020 sql.NullString `json:"02/06/2020"`
	Feb07_2020 sql.NullString `json:"02/07/2020"`
	Feb08_2020 sql.NullString `json:"02/08/2020"`
	Feb09_2020 sql.NullString `json:"02/09/2020"`
	Feb10_2020 sql.NullString `json:"02/10/2020"`
	Feb11_2020 sql.NullString `json:"02/11/2020"`
	Feb12_2020 sql.NullString `json:"02/12/2020"`
	Feb13_2020 sql.NullString `json:"02/13/2020"`
	Feb14_2020 sql.NullString `json:"02/14/2020"`
	Feb15_2020 sql.NullString `json:"02/15/2020"`
	Feb16_2020 sql.NullString `json:"02/16/2020"`
	Feb17_2020 sql.NullString `json:"02/17/2020"`
	Feb18_2020 sql.NullString `json:"02/18/2020"`
	Feb19_2020 sql.NullString `json:"02/19/2020"`
	Feb20_2020 sql.NullString `json:"02/20/2020"`
	Feb21_2020 sql.NullString `json:"02/21/2020"`
	Feb22_2020 sql.NullString `json:"02/22/2020"`
	Feb23_2020 sql.NullString `json:"02/23/2020"`
	Feb24_2020 sql.NullString `json:"02/24/2020"`
	Feb25_2020 sql.NullString `json:"02/25/2020"`
	Feb26_2020 sql.NullString `json:"02/26/2020"`
	Feb27_2020 sql.NullString `json:"02/27/2020"`
	Feb28_2020 sql.NullString `json:"02/28/2020"`
	Feb29_2020 sql.NullString `json:"02/29/2020"`
	Mar01_2020 sql.NullString `json:"03/01/2020"`
	Mar02_2020 sql.NullString `json:"03/02/2020"`
	Mar03_2020 sql.NullString `json:"03/03/2020"`
}

type DataWHO struct {
	ID         int            `json:"id"`
	State      sql.NullString `json:"state"`
	Country    sql.NullString `json:"country"`
	Who_region sql.NullString `json:"who_region"`
	Jan21_2020 sql.NullString `json:"01/21/2020`
	Jan22_2020 sql.NullString `json:"01/22/2020"`
	Jan23_2020 sql.NullString `json:"01/23/2020"`
	Jan24_2020 sql.NullString `json:"01/24/2020"`
	Jan25_2020 sql.NullString `json:"01/25/2020"`
	Jan26_2020 sql.NullString `json:"01/26/2020"`
	Jan27_2020 sql.NullString `json:"01/27/2020"`
	Jan28_2020 sql.NullString `json:"01/28/2020"`
	Jan29_2020 sql.NullString `json:"01/29/2020"`
	Jan30_2020 sql.NullString `json:"01/30/2020"`
	Jan31_2020 sql.NullString `json:"01/31/2020"`
	Feb01_2020 sql.NullString `json:"02/01/2020"`
	Feb02_2020 sql.NullString `json:"02/02/2020"`
	Feb03_2020 sql.NullString `json:"02/03/2020"`
	Feb04_2020 sql.NullString `json:"02/04/2020"`
	Feb05_2020 sql.NullString `json:"02/05/2020"`
	Feb06_2020 sql.NullString `json:"02/06/2020"`
	Feb07_2020 sql.NullString `json:"02/07/2020"`
	Feb08_2020 sql.NullString `json:"02/08/2020"`
	Feb09_2020 sql.NullString `json:"02/09/2020"`
	Feb10_2020 sql.NullString `json:"02/10/2020"`
	Feb11_2020 sql.NullString `json:"02/11/2020"`
	Feb12_2020 sql.NullString `json:"02/12/2020"`
	Feb13_2020 sql.NullString `json:"02/13/2020"`
	Feb14_2020 sql.NullString `json:"02/14/2020"`
	Feb15_2020 sql.NullString `json:"02/15/2020"`
	Feb16_2020 sql.NullString `json:"02/16/2020"`
	Feb17_2020 sql.NullString `json:"02/17/2020"`
	Feb18_2020 sql.NullString `json:"02/18/2020"`
	Feb19_2020 sql.NullString `json:"02/19/2020"`
	Feb20_2020 sql.NullString `json:"02/20/2020"`
	Feb21_2020 sql.NullString `json:"02/21/2020"`
	Feb22_2020 sql.NullString `json:"02/22/2020"`
	Feb23_2020 sql.NullString `json:"02/23/2020"`
	Feb24_2020 sql.NullString `json:"02/24/2020"`
}
