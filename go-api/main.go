package main

import (
	"net/http"
	"time"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Messages struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Message   string
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/messages", func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:@/golang_test")
		if err != nil {
			panic(err)
		}
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(148)
		db.SetMaxIdleConns(1)
		defer db.Close()

		insert, err := db.Query("INSERT INTO `messages` (`message`) VALUES ('Hello posting from Golang')")

		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()

		data := map[string]interface{}{
			"message": "Test posting from golang",
		}
		c.JSON(http.StatusOK, data)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
