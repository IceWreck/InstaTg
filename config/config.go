package config

import "flag"

const Version = "1.0.0"

// Config struct to hold all the configuration settings for our application.
type Config struct {
	InstagramUsername string
	InstagramPassword string
	TelegramBotToken  string
	TelegramChannel   int64
	DatabasePath      string // boltDB location
	TempMediaPath     string // temporary photo/video storage location
}

// GetConfig from command line flags
func GetConfig() Config {
	var cfg Config

	flag.StringVar(&cfg.InstagramUsername, "iguser", "", "Instagram Username")
	flag.StringVar(&cfg.InstagramPassword, "igpass", "", "Instagram Password")
	flag.StringVar(&cfg.TelegramBotToken, "tgtoken", "", "Telegram Bot Token")
	flag.Int64Var(&cfg.TelegramChannel, "tgchannel", 0, "Telegram Channel ID")

	flag.Parse()

	return cfg
}
