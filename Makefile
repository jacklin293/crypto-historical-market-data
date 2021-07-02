download:
	$(eval folder := "data/original-files/binance-${PAIR}-${INTERVAL}-kline-${YEAR}")
	mkdir -p ${folder}
	l='01 02 03 04 05 06 07 08 09 10 11 12'; for k in $$l; do wget "https://data.binance.vision/data/spot/monthly/klines/${PAIR}/${INTERVAL}/${PAIR}-${INTERVAL}-${YEAR}-$${k}.zip" -P ${folder}; done
unzip:
	$(eval source := "data/original-files/binance-${PAIR}-${INTERVAL}-kline-${YEAR}")
	$(eval dest := "data/csv/binance-${PAIR}-${INTERVAL}-kline-${YEAR}")
	mkdir -p ${dest}
	$(eval files := $(shell ls ${source}))
	@for k in $(files); do unzip "${source}/$${k}" -d ${dest}; done
