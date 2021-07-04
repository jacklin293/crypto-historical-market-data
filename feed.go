package main

import (
	"bufio"
	"crypto-historical-market-data/cryptodb"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func feedCsvToDB(db *cryptodb.DB, pair string, interval string, year string, month string) (err error) {
	// Create kline table
	tableName := cryptodb.GetTableName(pair)
	fmt.Printf(" - Create table '%s' if not exists\n", tableName)
	err = db.CreateBinanceBtcusdtKlineTableIfNotExists(tableName)
	if err != nil {
		return
	}

	// Reading and inserting in batch asynchronously
	if month == "all" {
		for i := 1; i <= 12; i++ {
			if err = handleReadAndInsert(db, pair, interval, year, strconv.Itoa(i)); err != nil {
				return
			}
		}
	} else {
		if err = handleReadAndInsert(db, pair, interval, year, month); err != nil {
			return
		}
	}

	return nil
}

func handleReadAndInsert(db *cryptodb.DB, pair string, interval string, year string, month string) (err error) {
	var lineCounter int
	lineCh := make(chan string, DB_KLINES_BATCH_INSERT_NUMBER)
	csvFilePath := getCsvFilePath(pair, interval, year, month)
	csvFileName := getCsvFileName(pair, interval, year, month)

	fmt.Printf(" - Processing %s\n", csvFileName)
	go readCsvFile(csvFilePath, DB_KLINES_BATCH_INSERT_NUMBER, lineCh, &lineCounter)
	if err = insertIntoDB(db, interval, lineCh, &lineCounter); err != nil {
		return
	}

	return
}

func readCsvFile(path string, batchNum int64, ch chan string, lineCounter *int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		lines = append(lines, string(line))
		if len(lines) == int(batchNum) {
			*lineCounter += len(lines)
			for i := 0; i < len(lines); i++ {
				ch <- lines[i]
			}
			lines = lines[:0] // clear the container
		}
	}

	// Send the rest of the lines
	*lineCounter += len(lines)
	for i := 0; i < len(lines); i++ {
		ch <- lines[i]
	}
}

func insertIntoDB(db *cryptodb.DB, interval string, lineCh chan string, lineCounter *int) (err error) {
	var lineInsertedCounter int
	var klines []cryptodb.BinanceBtcusdtKline
	for {
		select {
		case line := <-lineCh:
			klineData, err := processKlineRow(line)
			if err != nil {
				return err
			}
			kline := cryptodb.NewBinanceBtcusdtKline(interval, klineData)
			klines = append(klines, kline)

			// Batch insert into klines table
			if len(klines) == DB_KLINES_BATCH_INSERT_NUMBER {
				_, err := db.BatchInsertKlines(klines)
				if err != nil {
					if cryptodb.IsErrDupEntry(err) {
						fmt.Println("   -", err)
					} else {
						return err
					}
				} else {
					fmt.Printf("   - %d rows have been inserted successfully\n", len(klines))
				}

				// Reset the slice
				klines = klines[:0]

				lineInsertedCounter += DB_KLINES_BATCH_INSERT_NUMBER
			}

			// Deal with the rest of liens as the line number of csv could not be exactly a multiple of DB_KLINES_BATCH_INSERT_NUMBER
			if (len(klines) + lineInsertedCounter) == *lineCounter {
				_, err := db.BatchInsertKlines(klines)
				if err != nil {
					if cryptodb.IsErrDupEntry(err) {
						fmt.Println("   -", err)
					} else {
						return err
					}
				} else {
					fmt.Printf("   - %d rows have been inserted successfully\n", len(klines))
				}
				return nil
			}
		}
	}
	return nil
}

func processKlineRow(row string) (kline map[string]interface{}, err error) {
	cols := strings.Split(row, ",")

	openTimestamp, err := strconv.ParseInt(cols[0], 10, 64) // milliseconds
	if err != nil {
		return
	}
	openTime := time.Unix(0, openTimestamp*int64(time.Millisecond)).In(time.UTC)

	openPrice, err := strconv.ParseFloat(cols[1], 64)
	if err != nil {
		return
	}

	highPrice, err := strconv.ParseFloat(cols[2], 64)
	if err != nil {
		return
	}

	lowPrice, err := strconv.ParseFloat(cols[3], 64)
	if err != nil {
		return
	}

	closePrice, err := strconv.ParseFloat(cols[4], 64)
	if err != nil {
		return
	}

	volume, err := strconv.ParseFloat(cols[5], 64)
	if err != nil {
		return
	}

	closeTimestamp, err := strconv.ParseInt(cols[6], 10, 64) // milliseconds
	if err != nil {
		return
	}
	closeTime := time.Unix(0, closeTimestamp*int64(time.Millisecond)).In(time.UTC)

	kline = map[string]interface{}{
		"open":       openPrice,
		"high":       highPrice,
		"low":        lowPrice,
		"close":      closePrice,
		"volume":     volume,
		"open_time":  openTime,
		"close_time": closeTime,
	}
	return
}