package main

import (
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	bot, err := linebot.New(
		os.Getenv("06843eb2a88693584fb809abce4ecc88"),
		os.Getenv("NulDkgidRFKdTlMnZX3Ydm2FV+F5DyT+x2EhD731wXi6iQJaliF32JuLDDIcHzp1paDqoMoYKzJ4hSNk7V0qXxmxCQNpHqJ6Ouro+0hmJR6Z2hMp+I7+3SQoSvl92Pk7QtUblb5diALv1IgFlT3u5wdB04t89/1O/w1cDnyilFU="),
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

	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("port error"err)
	}

	/*
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("Could not start server. %s", err)
	}
	*/
}