package main

import (
	"log"
	"net/http"
	"os"
	"fmt"

	"github.com/line/line-bot-sdk-go/linebot"
)
/*
func pushMessage(bot Client, userID string) {
		if _, err := bot.PushMessage(userID, linebot.NewTextMessage("やっほ")).Do(); err != nil {
		log.Fatal(err)
	}
}
*/
func main() {
	port := os.Getenv("PORT")
	//userID := "Ubc0a1608b57a68e8fd8ec1c87fdc7697"
	if port == "" {
		port = "8080"
	}

	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Push Message
	//pushMessage(bot, userID)
	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}

		for _, event := range events {
			fmt.Printf(event.Source.UserID)
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					fmt.Println("Send", message.Text)
					// なあちゃんと言われた時はなあに？と返して
					// それ以外はAPIによる返事を返す
					if message.Text == "なあちゃん" {
						repmes := "なあに？"
						fmt.Println("Reply", repmes)
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(repmes)).Do(); err != nil {
							log.Print(err)
						}
					} else {
					repmes, err := a3rt(message.Text)
					if err != nil {
						log.Print(err)
					}
					fmt.Println("Reply", repmes.Results[0].Reply)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(repmes.Results[0].Reply)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	}
})

	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("port error", err)
	}
}