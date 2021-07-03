package cryptodb

import (
	"time"
)

type BinanceBTCUSDTKline struct {
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

func (db *DB) InsertBatchKlines() (id int64, err error) {

	/*
		k := BinanceBTCUSDTKline{
		}

		result := db.Create(&k)

		if user.ID == 0 || result.Error != nil || result.RowsAffected == 0 {

		}
		return user.ID, result.Error
	*/
	return 0, nil
}
