package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	FOLDER_DOWNLOAD         = "data/download"
	FOLDER_CSV              = "data/csv"
	BINANCE_PUBLIC_DATA_URL = "https://data.binance.vision/data/spot/monthly/klines"
	DB_DSN                  = "root:root@tcp(127.0.0.1:3306)/crypto_db?charset=utf8mb4"
	DB_KLINES_BATCH_NUMBER  = 200
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Pair
	fmt.Printf("Please enter a pair (e.g. BTCUSDT): ")
	pair, err := reader.ReadString('\n')
	pair = strings.Replace(pair, "\n", "", -1) // replace new line to empty from the input
	if err != nil {
		log.Fatal(err)
	}

	// Interval
	fmt.Printf("Please enter an interval (e.g. 1m, 1h, 1d, 1w, 1mo): ")
	interval, err := reader.ReadString('\n')
	interval = strings.Replace(interval, "\n", "", -1)
	if err != nil {
		log.Fatal(err)
	}

	// Year
	fmt.Printf("Please enter a year (e.g. 2021): ")
	year, err := reader.ReadString('\n')
	year = strings.Replace(year, "\n", "", -1)
	if err != nil {
		log.Fatal(err)
	}

	// Month
	fmt.Printf("Please enter a month (e.g. 1, 7, 12 or all): ")
	month, err := reader.ReadString('\n')
	month = strings.Replace(month, "\n", "", -1)
	if err != nil {
		log.Fatal(err)
	}

	// Download files
	fmt.Println("\nStart to download file(s)")
	if err = downloadFiles(pair, interval, year, month); err != nil {
		log.Fatal(err)
	}

	// Unzip files
	fmt.Println("\nStart to unzip file(s)")
	if err = unzipFiles(pair, interval, year, month); err != nil {
		log.Fatal(err)
	}

	// Feed csv into DB
	// fmt.Println("\nFeed data into database")
	// if err = feedCsvToDB(pair, interval, year, month); err != nil {
	//	log.Fatal(err)
	// }
}
