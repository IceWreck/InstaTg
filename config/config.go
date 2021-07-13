package config

import "flag"

const Version = "1.0.0"

// Config struct to hold all the configuration settings for our application.
type Config struct {
	InstagramUsername string
	InstagramPassword string
	InstagramChannel  string
	TelegramBotToken  string
	TelegramChannel   int64
	DatabasePath      string // boltDB location
	TempMediaPath     string // temporary photo/video storage location
	ReRunInterval     int64  // in minutes
}

// GetConfig from command line flags
func GetConfig() Config {
	var cfg Config

	flag.StringVar(&cfg.InstagramUsername, "iguser", "", "Your Instagram Username")
	flag.StringVar(&cfg.InstagramPassword, "igpass", "", "Your Instagram Password")
	flag.StringVar(&cfg.InstagramChannel, "igchan", "", "Instagram Channel's Username")
	flag.StringVar(&cfg.TelegramBotToken, "tgtoken", "", "Telegram Bot Token")
	flag.Int64Var(&cfg.TelegramChannel, "tgchannel", 0, "Telegram Channel ID")
	flag.StringVar(&cfg.DatabasePath, "dbpath", "./store.boltdb", "Database File Path (optional)")
	flag.Int64Var(&cfg.ReRunInterval, "rerun", 15, "Interval (in minutes) after which the bot shoudl check IG for more posts")

	flag.Parse()

	cfg.TempMediaPath = "./tmpdl"

	return cfg
}
