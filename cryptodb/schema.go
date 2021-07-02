package cryptodb

import (
	"bytes"
	"html/template"
	"strings"
)

const (
	TEMPLATE_KLINE_TABLE_SQL = "cryptodb/kline-table-sql.tmpl"
)

// Create kline table
func (db *DB) CreateKlineTable(tableName string) error {
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
