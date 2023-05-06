package datam

import (
     "time"
)
 // 关系型数据库连接 数据结构
type Conn struct {
        Id string
	Name string
	Ip   string
	Port string
	User string
	Pass string
        Dbname string
	Dbtype string
}
type ConnList [10]Conn

// 可执行SQL 数据结构
type Sqlseq struct {
       Id string
       Sqlstr string
}

// 执行的SQL历史纪录 数据结构 
type Sqllog struct {
       id string
       Conn_Name string
       Sqlstr string
       Exec_time time.Time
}

// Redis 数据库连接数据结构
type Redisconn struct {
     Addr string
     Password string
     DB string
}
