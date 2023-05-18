package main

import (
	badger "github.com/dgraph-io/badger/v4"
	"github.com/gin-gonic/gin"
	"log"
        "net/http"
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

        r.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
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
