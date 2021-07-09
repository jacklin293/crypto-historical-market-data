CREATE TABLE IF NOT EXISTS `klines` (
  `pair_interval` varchar(15) NOT NULL COMMENT 'pair+interval e.g. btcusdt_1h',
  `open` decimal(18,8) UNSIGNED NOT NULL COMMENT 'Open price',
  `high` decimal(18,8) UNSIGNED NOT NULL COMMENT 'High price',
  `low` decimal(18,8) UNSIGNED NOT NULL COMMENT 'Low price',
  `close` decimal(18,8) UNSIGNED NOT NULL COMMENT 'Close price',
  `volume` decimal(18,8) UNSIGNED NOT NULL COMMENT 'Volume',
  `open_time` datetime NOT NULL COMMENT 'Open time',
  `close_time` datetime NOT NULL COMMENT 'Close time'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Klines';

ALTER TABLE `klines`
  ADD UNIQUE KEY IF NOT EXISTS `pair_interval_opentime` (`pair_interval`,`open_time`);
