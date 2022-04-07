package instagram

import (
	"github.com/IceWreck/InstaTG/config"
	"github.com/ahmdrz/goinsta/v2"
)

func ConnectionCheck(app *config.Application) error {
	user := app.InstaBot.Account.Username
	// random API call to check if everything is working
	_, err := app.InstaBot.Profiles.ByName("instagram")
	if err == nil {
		app.Logger.Println("Instagram authenticated. Username: ", user)
	} else {
		// Panic is intentional because we dont want it to try making requests if it fails.
		// And if connection loss is just temporary, then systemd will restart anyways.
		app.Logger.Panicln("Instagram authentication error: ", err)
	}
	return err
}

func NewConnection(app *config.Application) *goinsta.Instagram {
	insta := goinsta.New(app.Config.InstagramUsername, app.Config.InstagramPassword)
	err := insta.Login()
	if err != nil {
		app.Logger.Panicln("Instagram API connection error: ", err)
	}
	return insta
}
