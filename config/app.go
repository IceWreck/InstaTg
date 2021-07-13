package config

import (
	"log"

	"github.com/ahmdrz/goinsta/v2"
	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Application struct to hold the dependencies for our application.
type Application struct {
	Config   Config
	Logger   *log.Logger
	TGBot    *tgbotapi.BotAPI
	InstaBot *goinsta.Instagram
	BoltDB   *bolt.DB
}
