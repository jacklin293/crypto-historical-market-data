# crypto-historical-market-data

This is a tool to transfer historical `kline` (or `candlestick`) data provided by
[Binance public data](https://github.com/binance/binance-public-data/) into MySQL database.

This project will do these for you:

* Set up MySQL database
* Download and decompress files from Binance public data
* Transfer historical data (csv) to database
    * Note: Only these fields will be stored into database
        * `interval`, `open`, `high`, `low`, `close`, `volume`, `open time`, `close time`

Once it's all set, you can start to use it whatever you prefer.

# Requirements

* Go version (for transfer program): `1.16.5` above (can't gurantee any version below it)

# Set up the environment

1. Set up database

```
docker-compose up
```

> Folder `mariadb` will be created, so that you won't miss any data when db container is killed.

2. Check if database is up

> You might want to change ports in `docker-compose.yml` if ports collide.

Connect to database directly

```
mysql -h 127.0.0.1 -u root -proot
```

Or use phpmyadmin (if no error is shown on the page, it means fine)

```
http://localhost:8080/
```

3. Install golang: https://golang.org/doc/install

After installation, please execute below in this project:

```
go mod download
go build
```

# How to use

Execute the binary file with params

```
./crypto-historical-market-data -pair=BTCUSDT -interval=1h -year=2021 -month=1
```

Execute the binary file without params, then you will be asked to fill them

```
./crypto-historical-market-data

Please enter a pair (e.g. BTCUSDT): BTCUSDT
Please enter an interval (e.g. 1m, 1h, 1d, 1w, 1mo): 1h
Please enter a year (e.g. 2021): 2021
Please enter a month (e.g. 1, 7, 12 or all): 1
```

> INTERVAL: `1m`, `3m`, `5m`, `15m`, `30m`, `1h`, `2h`, `4h`, `6h`, `8h`, `12h`, `1d`, `3d`, `1w`, `1mo`



