package main

import (
	models "facebook/goapi"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	r := gin.Default()

	db, err := models.CreateConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	readUser, err := models.GetAjay(db, 1)
	if err != nil {
		log.Fatal(err)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, readUser)
	})

	r.GET("/all", func(c *gin.Context) {
		rows, err := models.FetchAllRows(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, rows)
	})

	r.Run()
}
