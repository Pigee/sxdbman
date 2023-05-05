package main

import (
	"fmt"
	badger "github.com/dgraph-io/badger/v4"
	"log"
	"sxdbman/datam"
)

func main() {

	fmt.Println("Hello,Man")
	// https://blog.csdn.net/u013317445/article/details/114841853 struct slice
	myconn := datam.Conn{"ljp", "fdfd", "fdfd", "fdfdf", "fdfdf"}
	myconn.ConnPort = "2323"
	mcl := make([]datam.Conn, 3)
	//append(mcl,myconn)
	mcl[0] = myconn
	mcl = append(mcl, datam.Conn{"ljp2", "fdfd", "fdfd", "fdfdf", "fdfdf"})

	fmt.Println(myconn)
	fmt.Println(mcl)

	db, err := badger.Open(badger.DefaultOptions("./resources/badger"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
