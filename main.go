package main

// TODO check the data source doesn't miss any point (every 4 hour)
// TODO throw an error when anything went wrong
// TODO Make `interval` as a column
// TODO make this a command to enter specific data

import (
	"crypto-historical-market-data/cryptodb"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	FOLDER_CSV = "data/csv"
	DB_DSN     = "root:root@tcp(127.0.0.1:3306)/crypto_db?charset=utf8mb4"
)

func main() {
	listFolder()

	// TODO Create folder
	// TODO Downlaod files
	// TODO Unzip files
	// TODO Feed csv into DB
}

func listFolder() {
	files, err := ioutil.ReadDir(FOLDER_CSV)
	if err != nil {
		log.Fatal(err)
	}
	db, err := cryptodb.NewDB(DB_DSN)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		// Split 'binance-BTCUSDT-4h-kline-2021' into a slice
		s := strings.Split(f.Name(), "-")
		tableName := fmt.Sprintf("%s_%s_klines", s[0], s[1])
		if err := db.CreateKlineTable(tableName); err != nil {
			log.Fatal(err)
		}
		log.Printf("Table `%s` was created successfully", tableName)
	}
}
