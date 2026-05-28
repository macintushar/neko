package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"log"

	"github.com/joho/godotenv"

	"github.com/thecodearcher/limen"
	credentialpassword "github.com/thecodearcher/limen/plugins/credential-password"

	gormadapter "github.com/thecodearcher/limen/adapters/gorm"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	port, portExists := os.LookupEnv("PORT")

	if !portExists {
		log.Println("PORT doesn't exist. Defaulting to :3040")
		port = "3080"
	} else {
		log.Println("port", port)
	}

	url, urlExists := os.LookupEnv("BASE_URL")

	if !urlExists {
		log.Println("BASE_URL is not set. Defaulting to 0.0.0.0")
	}

	serverUrl := url + ":" + port

	db, err := gorm.Open(sqlite.Open("neko.db"))

	if err != nil {
		log.Fatal(err)
		panic("failed to connect database")
	}

	config := &limen.Config{
		BaseURL:  "http://localhost:3040",
		Database: gormadapter.New(db),
		Plugins: []limen.Plugin{
			credentialpassword.New(),
		},
	}

	auth, err := limen.New(config)

	if err != nil {
		log.Fatalf("Failed to create limen: %v", err)
	}

	handler := auth.Handler()

	router := gin.Default()

	router.Any("/auth/*path", func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(serverUrl)
}
