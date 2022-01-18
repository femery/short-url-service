# 短链接服务

```sql
CREATE TABLE `short_urls`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT,
    `surl`       varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '短链标志',
    `lurl`       varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '原始链接',
    `ctime`      datetime NOT NULL                DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `last_vtime` datetime NOT NULL                DEFAULT CURRENT_TIMESTAMP COMMENT '最后访问时间',
    `times`      int(11) NOT NULL DEFAULT '0' COMMENT '访问次数',
    `status`     int(2) NOT NULL DEFAULT '0' COMMENT '0正常；1失效',
    PRIMARY KEY (`id`),
    UNIQUE KEY `surl` (`surl`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
```

// 关于统计 // 将一定时间内的变化量存储到容器中，然后启动一个定时任务捞出来，聚合存储到数据库
