package instagram

import (
	"github.com/IceWreck/InstaTG/config"
	"github.com/ahmdrz/goinsta/v2"
)

func ConnectionCheck(app *config.Application) error {
	user := app.InstaBot.Account.Username
	// random API call to check if everything is working
	_, err := app.InstaBot.Profiles.ByName("facebook")
	if err == nil {
		app.Logger.Println("Instagram authenticated. Username: ", user)
	} else {
		app.Logger.Println("Instagram authentication error: ", err)
	}
	return err
}

func NewConnection(app *config.Application) *goinsta.Instagram {
	insta := goinsta.New(app.Config.InstagramUsername, app.Config.InstagramPassword)
	err := insta.Login()
	if err != nil {
		app.Logger.Fatalln("Instagram API connection error: ", err)
	}
	return insta
}
