package datam

import ()

type Conn struct {
	ConnName string
	ConnIp   string
	ConnPort string
	ConnUser string
	ConnPass string
}

type ConnList []Conn
