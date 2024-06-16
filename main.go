package main

import (
	"log"

	"github.com/NessibeliY/telegram-bot/internal/config"
	"github.com/NessibeliY/telegram-bot/internal/values"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	filePath := values.ConfigFile

	cfg, err := config.Load(filePath)
	if err != nil {
		log.Printf("Error loading config: %v", err)
		return
	}

	bot, err := tgbotapi.NewBotAPI(cfg.TgToken)
	if err != nil {
		log.Printf("Error creating bot: %v", err)
		return
	}

	// Set the bot to use debug mode (verbose logging).
	bot.Debug = true
	log.Printf("Authorized as @%s", &bot.Self.UserName)

	// Set up updates configuration.
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Get updates from the Telegram API.
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Printf("Error getting updates: %v", err)
	}

	// Process incoming messages.
	for update := range updates {
		if update.Message == nil { // Ignore any non-Message updates.
			continue
		}

		// Print received message text and sender username.
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// Respond to the user.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello! I am your Telegram bot.")
		_, err := bot.Send(msg)
		if err != nil {
			log.Printf("Error sending message: %v", err)
		}
	}
}
