# crypto-historical-market-data

This is a tool to transfer historical `kline` (or `candlestick`) data provided by
[Binance public data](https://github.com/binance/binance-public-data/) into MySQL database.

This project will do these for you:

* Set up MySQL database
* Download and `unzip` files from Binance public data
* Transfer historical data (csv) to database

Once it's all set, you can start to use it whatever you prefer.

# Requirements

* Commands (for Makefile): `make` `wget` `unzip`
* Go version (for transfer program): `1.16.5` above (can't gurantee any version below it)

# Set up the environment

1. Set up database

```
docker-compose up
```

> Folder `mariadb` will be created, so that you won't miss any data when db container is killed.

2. Check if database is up

> You might want to change ports in `docker-compose.yml` if it collides.

Connect to database directly

```
mysql -h 127.0.0.1 -u root -proot
```

Or use phpmyadmin (if no error is shown on the page, it means fine)

```
http://localhost:8080/
```

# Feed data into database

1. Download klines data

Download klines data for a specific year, e.g.:

```
make download PAIR=BTCUSDT INTERVAL=4h PAIR=BTCUSDT YEAR=2021
```

> INTERVAL: `1m`, `3m`, `5m`, `15m`, `30m`, `1h`, `2h`, `4h`, `6h`, `8h`, `12h`, `1d`, `3d`, `1w`, `1mo`

Unzip the source data

```
make unzip PAIR=BTCUSDT INTERVAL=4h PAIR=BTCUSDT YEAR=2021
```



2. Transfer data into database

TODO


