package main

import (
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	bot, err := linebot.New(
		os.Getenv("fe5a2cbeb9c5cfa20434ff6691ddc642"),
		os.Getenv("1Jm0jNA9ovgwCyD9MMEQT1P8R+nsaf+uzPiZKF0NcHbGlS6qOQWC7Xzszy4d1q2MpaDqoMoYKzJ4hSNk7V0qXxmxCQNpHqJ6Ouro+0hmJR4pvFp6fqnPEpIW4zfMclRFHeVK/WHE1bja+XVrxIq+RAdB04t89/1O/w1cDnyilFU="),
	)
	if err != nil {
		log.Fatal(err)
	}

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
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	/*
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	*/

}