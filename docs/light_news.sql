CREATE DATABASE IF NOT EXISTS `light_news`;

USE `light_news`;

CREATE TABLE `news_model` (
  `news_id` int(11) NOT NULL AUTO_INCREMENT,
  `news_url` varchar(256) NOT NULL,
  `title` varchar(64) NOT NULL DEFAULT '',
  `rank` tinyint(4) NOT NULL DEFAULT '0',
  `author` varchar(32) NOT NULL DEFAULT '',
  `abstract` varchar(512) NOT NULL DEFAULT '' COMMENT '摘要',
  `publish_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `is_hot` tinyint(1) NOT NULL DEFAULT '0',
  `img_url` varchar(256) NOT NULL DEFAULT '',
  `list_url` varchar(256) NOT NULL DEFAULT '',
  `raw_url` varchar(256) NOT NULL DEFAULT '',
  `data_source` varchar(32) NOT NULL DEFAULT '',
  PRIMARY KEY (`news_id`),
  UNIQUE KEY `uk_news_url` (`news_url`)
) ENGINE=InnoDB AUTO_INCREMENT=7999 DEFAULT CHARSET=utf8;