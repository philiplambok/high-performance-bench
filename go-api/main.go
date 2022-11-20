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
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.POST("/messages", func(c *gin.Context) {
		// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
		db, err := sql.Open("mysql", "root:@/golang_test")
		if err != nil {
			panic(err)
		}
		// See "Important settings" section.
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(148)
		db.SetMaxIdleConns(1)
		defer db.Close()

		// perform a db.Query insert
		insert, err := db.Query("INSERT INTO `messages` (`message`) VALUES ('Hello posting from Golang')")

		// if there is an error inserting, handle it
		if err != nil {
			panic(err.Error())
		}
		// be careful deferring Queries if you are using transactions
		defer insert.Close()

		data := map[string]interface{}{
			"message": "Test posting from golang",
		}
		c.JSON(http.StatusOK, data)
	})
	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
