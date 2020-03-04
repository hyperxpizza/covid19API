import sqlite3
import pandas as pd
import os

dirpath = os.path.dirname(os.path.abspath(__file__))

csvfiles = {
    "csse_confirmed": 'csv/time_series_19-covid-Confirmed.csv',
    "csse_deaths": 'csv/time_series_19-covid-Deaths.csv',
    "csse_recovered": 'csv/time_series_19-covid-Recovered.csv',
    "who": 'csv/who_covid_19_sit_rep_time_series.csv'
}



for table, filename in csvfiles.items():
    # load data
    df = pd.read_csv(dirpath + "/" + filename)

    # strip whitespace from headers
    df.columns = df.columns.str.strip()

    con = sqlite3.connect("../covid19.db")

    # drop data into database
    df.to_sql(table, con)

con.close()