
-- 图书表
CREATE TABLE `books` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  `author` varchar(16) NOT NULL DEFAULT '',
  `content` text NOT NULL,
  `status` tinyint(4) unsigned NOT NULL DEFAULT '0',
  `creator` varchar(16) NOT NULL DEFAULT '',
  `created_at` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
