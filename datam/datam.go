package datam

import ()

type Conn struct {
	Name string
	Ip   string
	Port string
	User string
	Pass string
        Dbname string
	Dbtype string
}

type ConnList [10]Conn
