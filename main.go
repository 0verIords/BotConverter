package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rylio/ytdl"
)

func youtubeDown(url string) {
	vid, _ := ytdl.GetVideoInfo(url)
	file, _ := os.Create(vid.ID + ".mp3")

	defer file.Close()

	vid.Download(vid.Formats[17], file)

}

func Video_id(url string) (string, bool) {
	if strings.Contains(url, "=") {
		youtubeDown(url)
		url := strings.Split(url, "=")
		return url[len(url)-1], true
	} else if strings.Contains(url, "/") {

		youtubeDown(url)
		url := strings.Split(url, "/")
		return url[len(url)-1], true

	} else {
		return "fail", false
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI(giveKey("token"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		{
			if update.Message == nil { // ignore any non-Message Updates
				continue
			}

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
