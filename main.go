package main

import (
  "gopkg.in/telegram-bot-api.v4"
  "log"
  "net/http"
  "os"
)

func main() {
  port := os.Getenv("WEBHOOK_PORT")
  telegram_token := os.Getenv("TELEGRAM_TOKEN")
  telegram_server := os.Getenv("TELEGRAM_SERVER")


  bot, err := tgbotapi.NewBotAPI( telegram_token )
  if err != nil {
    log.Panic(err)
  }

  //////////////////////////////////////////////////////////////
  // bot.Debug = true

  // log.Printf("Authorized on account %s", bot.Self.UserName)

  // u := tgbotapi.NewUpdate(0)
  // u.Timeout = 60

  // updates, err := bot.GetUpdatesChan(u)

  // for update := range updates {
  //   if update.Message == nil {
  //     continue
  //   }

  //   log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

  //   reply_content := ""
  //   switch update.Message.Text {
  //   case "Hello":
  //     reply_content = "Hi ar!"
  //   case "Bye":
  //     reply_content = "Dun say goodbye la."
  //   default:
  //     reply_content = "You up what?"
  //   }

  //   msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply_content)
  //   msg.ReplyToMessageID = update.Message.MessageID

  //   bot.Send(msg)
  // }
  //////////////////////////////////////////////////////////////


  webhook_url := telegram_server + ":" + port + "/"+ bot.Token

  bot.Debug = true

  log.Printf("Authorized on account %s", bot.Self.UserName)
  log.Printf("Webhook URL %s", webhook_url)
  
  resp, err := http.Get("https://api.telegram.org/bot"+bot.Token+"/setWebhook?url=")
  if err != nil {
    log.Fatal(err)
  }
  
  _, err = bot.SetWebhook(tgbotapi.NewWebhook(webhook_url) )
  if err != nil {
    log.Fatal(err)
  }

  updates := bot.ListenForWebhook("/" + bot.Token)
  go http.ListenAndServe(":" + port, nil)

  for update := range updates {
    if update.Message == nil {
      continue
    }

    log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

    reply_content := ""
    switch update.Message.Text {
    case "Hello":
      reply_content = "Hi ar!"
    case "Bye":
      reply_content = "Dun say goodbye la."
    default:
      reply_content = "You up what?"
    }

    msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply_content)
    msg.ReplyToMessageID = update.Message.MessageID

    bot.Send(msg)
  }

}