package http

import (
	"context"
	"github.com/kinnch/line-todo/core/todo"

	"github.com/kinnch/line-todo/models"
	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
)

type HTTPCallBackHanlder struct {
	Bot          *linebot.Client
	ServicesInfo *models.ServicesInfo
}

// NewServiceHTTPHandler provide the inititail set up service path to handle request
func NewServiceHTTPHandler(e *echo.Echo, linebot *linebot.Client, servicesInfo *models.ServicesInfo) {

	hanlders := &HTTPCallBackHanlder{Bot: linebot, ServicesInfo: servicesInfo}
	e.GET("/ping", func(c echo.Context) error {

		return c.String(200, "Line boi Service : We are good thank you for asking us.")
	})
	e.POST("/callback", hanlders.Callback)

	e.POST("/login", loginHandler)
}

// Callback provides the function to handle request from line
func (handler *HTTPCallBackHanlder) Callback(c echo.Context) error {

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	events, err := handler.Bot.ParseRequest(c.Request())
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.String(400, linebot.ErrInvalidSignature.Error())
		} else {
			c.String(500, "internal")
		}
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				messageFromPing := todo.TodoController(event.Source.UserID, message.Text)
				if _, err = handler.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(messageFromPing)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
	return c.JSON(200, "")
}

func loginHandler(c echo.Context) error {

}
