package main

import (
	"log"
	"os"

	"github.com/kinnch/line-todo/servicemanagement"
	"github.com/kinnch/line-todo/servicemanagement/delivery/http"

	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
)

func connectLineBot() *linebot.Client {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}
	return bot
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// router := gin.New()
	// router.Use(gin.Logger())
	// router.LoadHTMLGlob("templates/*.tmpl.html")
	// router.Static("/static", "static")

	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl.html", nil)
	// })

	// router.GET("/mark2", func(c *gin.Context) {
	// 	c.String(http.StatusOK, string(blackfriday.MarkdownBasic([]byte("**hi!5555**"))))
	// })

	// router.Run(":" + port)
	e := echo.New()
	bankCoreInfo := servicemanagement.NewBankCoreServiceInfo()
	http.NewServiceHTTPHandler(e, connectLineBot(), bankCoreInfo)
	e.Logger.Fatal(e.Start(":" + port))
}
