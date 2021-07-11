package cryptodb

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

const (
	TALBE_KLINES_SQL_PATH = "cryptodb/table-klines.sql"
)

type Kline struct {
	KlineKey  string // ma type+pair+interval e.g. btcusdt_1h
	Open      decimal.Decimal
	High      decimal.Decimal
	Low       decimal.Decimal
	Close     decimal.Decimal
	Volume    decimal.Decimal
	OpenTime  time.Time
	CloseTime time.Time
}

func getKlineKey(pair string, interval string) string {
	return fmt.Sprintf("%s_%s", strings.ToLower(pair), interval)
}

func NewKline(data map[string]interface{}) Kline {
	pair := strings.ToLower(data["pair"].(string))
	interval := data["interval"].(string)
	return Kline{
		KlineKey:  getKlineKey(pair, interval),
		Open:      data["open"].(decimal.Decimal),
		High:      data["high"].(decimal.Decimal),
		Low:       data["low"].(decimal.Decimal),
		Close:     data["close"].(decimal.Decimal),
		Volume:    data["volume"].(decimal.Decimal),
		OpenTime:  data["open_time"].(time.Time),
		CloseTime: data["close_time"].(time.Time),
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
