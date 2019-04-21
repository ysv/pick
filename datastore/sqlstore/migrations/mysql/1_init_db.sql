-- +migrate Up

CREATE TABLE `pageviews` (
  `id` varchar(31) NOT NULL,
  `hostname` varchar(255) NOT NULL,
  `pathname` varchar(255) NOT NULL,
  `is_new_visitor` tinyint(1) NOT NULL,
  `is_new_session` tinyint(1) NOT NULL,
  `is_unique` tinyint(1) NOT NULL,
  `is_bounce` tinyint(1) DEFAULT NULL,
  `referrer` varchar(255) DEFAULT NULL,
  `duration` int(4) DEFAULT NULL,
  `timestamp` datetime NOT NULL,
  `site_tracking_id` varchar(8) NOT NULL,
  `is_finished` tinyint(1) NOT NULL DEFAULT '0'
);

-- +migrate Down
DROP TABLE IF EXISTS pageviews;
