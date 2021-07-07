package cryptodb

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

const (
	TALBE_KLINES_SQL_PATH = "cryptodb/table-klines.sql"
)

type Kline struct {
	PairInterval string // pair+interval e.g. btcusdt_1h
	Open         float64
	High         float64
	Low          float64
	Close        float64
	Volume       float64
	OpenTime     time.Time
	CloseTime    time.Time
}

func NewKline(pair string, interval string, data map[string]interface{}) Kline {
	return Kline{
		PairInterval: fmt.Sprintf("%s_%s", strings.ToLower(pair), interval),
		Open:         data["open"].(float64),
		High:         data["high"].(float64),
		Low:          data["low"].(float64),
		Close:        data["close"].(float64),
		Volume:       data["volume"].(float64),
		OpenTime:     data["open_time"].(time.Time),
		CloseTime:    data["close_time"].(time.Time),
	}
}

// Create kline table
func (db *DB) CreateTableKlinesIfNotExists() error {
	body, err := ioutil.ReadFile(TALBE_KLINES_SQL_PATH)
	if err != nil {
		return err
	}

	// Separate SQL statements
	stats := strings.Split(string(body), "\n\n")
	for _, stat := range stats {
		if result := db.GormDB.Exec(stat); result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (db *DB) BatchInsertKlines(klines []Kline) (int64, error) {
	result := db.GormDB.Create(klines)
	return result.RowsAffected, result.Error
}
