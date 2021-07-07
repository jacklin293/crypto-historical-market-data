CREATE TABLE IF NOT EXISTS `klines` (
  `pair_interval` varchar(15) NOT NULL COMMENT 'pair+interval e.g. btcusdt_1h',
  `open` float UNSIGNED NOT NULL COMMENT 'Open price',
  `high` float UNSIGNED NOT NULL COMMENT 'High price',
  `low` float UNSIGNED NOT NULL COMMENT 'Low price',
  `close` float UNSIGNED NOT NULL COMMENT 'Close price',
  `volume` float UNSIGNED NOT NULL COMMENT 'Volume',
  `open_time` datetime NOT NULL COMMENT 'Open time',
  `close_time` datetime NOT NULL COMMENT 'Close time'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Klines';

ALTER TABLE `klines`
  ADD UNIQUE KEY IF NOT EXISTS `pair_interval_opentime` (`pair_interval`,`open_time`);
