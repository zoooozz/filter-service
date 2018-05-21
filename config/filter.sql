
-- 业务表

CREATE TABLE `business` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'ctime',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'mtime',
  `name` varchar(32) NOT NULL DEFAULT '' COMMENT '业务组',
  `flag` varchar(16) NOT NULL DEFAULT '' COMMENT '标识',
  `state` tinyint(4) NOT NULL COMMENT '状态',
  PRIMARY KEY (`id`),
  UNIQUE KEY `flag` (`flag`),
  KEY `ctime` (`ctime`),
  KEY `mtime` (`mtime`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='业务表';

-- 关键词表

CREATE TABLE `keyword` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '新增时间',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `content` varchar(100) NOT NULL DEFAULT '' COMMENT '关键词',
  PRIMARY KEY (`id`),
  KEY `ctime` (`ctime`),
  KEY `mtime` (`mtime`),
  KEY `content` (`content`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT '关键词表';

-- 关系表

CREATE TABLE `relation` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'ctime',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'mtime',
  `keyword_id` int(11) NOT NULL COMMENT '敏感词',
  `state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否启用',
  `level` tinyint(4) NOT NULL DEFAULT '1' COMMENT '敏感词级别',
  `flag` varchar(32) NOT NULL DEFAULT 'all' COMMENT '拥有组',
  PRIMARY KEY (`id`),
  KEY `ctime` (`ctime`),
  KEY `mtime` (`mtime`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8;