package main

import (
	models "facebook/goapi"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	router := gin.Default()

	db, err := models.CreateConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	readUser, err := models.GetAjay(db, 1)
	if err != nil {
		log.Fatal(err)
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, readUser)
	})

	router.GET("/all", func(c *gin.Context) {
		rows, err := models.FetchAllRows(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, rows)
	})

	router.POST("/", func(c *gin.Context) {
		var request models.Ajay
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		insertUser, err := models.InsertAjay(db, request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, insertUser)

	})

	router.PUT("/", func(c *gin.Context) {
		var request models.Ajay
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		insertUser, err := models.UpdateAjay(db, &request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, insertUser)

	})

	router.Run(":9090")
}
