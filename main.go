package main

import (
	badger "github.com/dgraph-io/badger/v4"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	ginping()
}

func ginping() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":9999") // list

}

func getbadgerdb() {
	db, err := badger.Open(badger.DefaultOptions("./resources/badger"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
