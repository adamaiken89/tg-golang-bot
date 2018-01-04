package main

import (
  "log"
  "os"
  "gopkg.in/telegram-bot-api.v4"
)

func main() {
  bot, err := tgbotapi.NewBotAPI( os.Getenv("TELEGRAM_TOKEN") )
  if err != nil {
    log.Panic(err)
  }

  bot.Debug = true

  log.Printf("Authorized on account %s", bot.Self.UserName)

  u := tgbotapi.NewUpdate(0)
  u.Timeout = 60

  updates, err := bot.GetUpdatesChan(u)

  for update := range updates {
    if update.Message == nil {
      continue
    }

    log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)


    switch update.Message.Text {
    case "Hello":
      reply_content := "Hi ar!"
    case "Bye":
      reply_content := "Dun say goodbye la."
    default:
      reply_content := "You up what?"
    }

    msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply_content)
    msg.ReplyToMessageID = update.Message.MessageID

    bot.Send(msg)
  }
}