package cryptodb

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
	"time"
)

const (
	TEMPLATE_KLINE_TABLE_SQL = "cryptodb/kline-table-sql.tmpl"
)

type BinanceBtcusdtKline struct {
	Id            int64
	KlineInterval string
	Open          float64
	High          float64
	Low           float64
	Close         float64
	Volume        float64
	OpenTime      time.Time
	CloseTime     time.Time
}

func GetTableName(pair string) string {
	return fmt.Sprintf("binance_%s_klines", strings.ToLower(pair))
}

func NewBinanceBtcusdtKline(interval string, data map[string]interface{}) BinanceBtcusdtKline {
	return BinanceBtcusdtKline{
		KlineInterval: interval,
		Open:          data["open"].(float64),
		High:          data["high"].(float64),
		Low:           data["low"].(float64),
		Close:         data["close"].(float64),
		Volume:        data["volume"].(float64),
		OpenTime:      data["open_time"].(time.Time),
		CloseTime:     data["close_time"].(time.Time),
	}
}

// Create kline table
func (db *DB) CreateBinanceBtcusdtKlineTableIfNotExists(tableName string) error {
	replacement := map[string]string{
		"table_name": tableName,
	}
	var buf bytes.Buffer
	klineTableSqlTemplate := template.Must(template.ParseFiles(TEMPLATE_KLINE_TABLE_SQL))
	err := klineTableSqlTemplate.Execute(&buf, replacement)
	if err != nil {
		return err
	}

	// The template contains a few sql statement
	stats := strings.Split(buf.String(), "\n\n")
	for _, stat := range stats {
		if result := db.GormDB.Exec(stat); result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (db *DB) BatchInsertKlines(klines []BinanceBtcusdtKline) (int64, error) {
	result := db.GormDB.Create(klines)
	return result.RowsAffected, result.Error
}
