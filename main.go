package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

// Setting applying
func init() {
	err := godotenv.Load("settings.env")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Fatal("ERROR: Can't init API connection: " + err.Error())
	}
	log.Println(bot.GetWebhookInfo())

	wh, _ := tgbotapi.NewWebhookWithCert(os.Getenv("MYURL")+"/"+bot.Token, nil)
	bot.Request(wh)

	updates := bot.ListenForWebhook("/" + bot.Token)

	go http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), nil)

	adminid, err := strconv.Atoi(os.Getenv("ADMINID"))
	if err != nil {
		log.Fatal("ERROR: Wrong AdminID: " + err.Error())
	}
	for update := range updates {
		if update.Message != nil {
			if update.Message.Text == "/start" {
				bot.Send(tgbotapi.NewMessage(update.Message.From.ID, os.Getenv("STARTTEXT")))
			} else {
				msg := tgbotapi.NewMessage(int64(adminid), fmt.Sprintf("New message from: [ %s %s ](tg://user?id=%d)", update.Message.From.FirstName, update.Message.From.LastName, update.Message.From.ID))
				msg.ParseMode = "MarkdownV2"
				_, err := bot.Send(msg)
				if err != nil {
					log.Print("ERROR: Can't copy message to admin: " + err.Error())
					continue
				}
				_, err = bot.Send(tgbotapi.NewCopyMessage(int64(adminid), update.Message.From.ID, update.Message.MessageID))
				if err != nil {
					log.Print("ERROR: Can't copy message to admin: " + err.Error())
					continue
				}
				bot.Send(tgbotapi.NewMessage(update.Message.From.ID, os.Getenv("DONETEXT")))
			}
		}
	}
}
