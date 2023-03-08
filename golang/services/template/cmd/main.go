package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func main() {
	// config
	port := os.Getenv("PORT")
	cfg := DBConfig{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		Name:     os.Getenv("MYSQL_DATABASE"),
	}

	// init database
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// migration (MEMO: マイグレーションは別でプロセスでやったほうがいい)
	type Ping struct {
		gorm.Model
		Message string
	}
	if err := db.AutoMigrate(&Ping{}); err != nil {
		log.Fatal(err)
	}

	// init http server
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		ping := &Ping{Message: "pong"}
		db.Create(ping)
		c.JSON(http.StatusOK, gin.H{
			"message":  ping.Message,
			"RDMS":     db.Name(),
			"database": db.Migrator().CurrentDatabase(),
		})
	})

	log.Fatal(r.Run(fmt.Sprintf(":%s", port)))
}
