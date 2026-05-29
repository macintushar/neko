package main

import (
	"io/fs"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"embed"
	"log"

	"github.com/joho/godotenv"

	"github.com/thecodearcher/limen"
	credentialpassword "github.com/thecodearcher/limen/plugins/credential-password"

	gormadapter "github.com/thecodearcher/limen/adapters/gorm"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//go:embed all:ui-dist
var staticFS embed.FS

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Print("No .env file found. Using defaults", err)
	}

	port, portExists := os.LookupEnv("PORT")

	if !portExists {
		log.Println("PORT is not set. Defaulting to port :3090")
		port = "3090"
	} else {
		log.Println("PORT :", port)
	}

	url, urlExists := os.LookupEnv("BASE_URL")

	if !urlExists {
		log.Println("BASE_URL is not set. Defaulting to 0.0.0.0")
		url = "0.0.0.0"
	}

	serverURL := url + ":" + port

	db, err := gorm.Open(sqlite.Open("neko.db"))

	if err != nil {
		log.Fatal(err)
		panic("failed to connect database")
	}

	config := &limen.Config{
		BaseURL:  serverURL,
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

	uiDist, _ := fs.Sub(staticFS, "ui-dist")

	router.GET("/app", func(c *gin.Context) {
		data, _ := staticFS.ReadFile("ui-dist/index.html")
		c.Data(200, "text/html; charset=utf-8", data)
	})

	router.StaticFS("/app", http.FS(uiDist))

	router.Any("/auth/*path", func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	})

	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/app")
	})

	router.GET("/api/v1/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(serverURL)
}
