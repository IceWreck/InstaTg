package main

import (
	"log"
	"os"

	"github.com/IceWreck/InstaTG/config"
	"github.com/IceWreck/InstaTG/telegram"
)

func main() {
	app := &config.Application{
		Config: config.GetConfig(),
		Logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
	}

	app.Logger.Println("Hey !")
	app.TGBot = telegram.NewConnection(app)

	telegram.ConnectionCheck(app)
	// telegram.SendTextMessage(app, "Heya !")
	telegram.SendVideo(app, "./samples/like.mp4", "We only get one like.")
	// telegram.SendImage(app, "./samples/tux.jpg", "Here Lies Tux !")

}
