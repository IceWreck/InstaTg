// Fetch the latest posts from an Instagram channel and
// send them to a Telegram channel.
// It keeps track of sent posts so you dont have to worry about repeat stuff.
// Run this as a system service as it fetches new posts after <your set interval>.

package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/IceWreck/InstaTG/config"
	"github.com/IceWreck/InstaTG/instagram"
	"github.com/IceWreck/InstaTG/store"
	"github.com/IceWreck/InstaTG/telegram"
)

func main() {
	app := &config.Application{
		Config: config.GetConfig(),
		Logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
	}

	app.Logger.Println("Starting InstaTg. Version:", config.Version)
	app.Logger.Println("Your Instagram User:", app.Config.InstagramUsername)
	app.Logger.Println("Instagram Channel:", app.Config.InstagramChannel)

	app.Logger.Println("Connecting to Database ...")
	app.BoltDB = store.NewConnection(app)
	defer app.BoltDB.Close()

	app.Logger.Println("Connecting to Telegram ...")
	app.TGBot = telegram.NewConnection(app)
	telegram.ConnectionCheck(app)

	app.Logger.Println("Connecting to Instagram ...")
	app.InstaBot = instagram.NewConnection(app)
	instagram.ConnectionCheck(app)

	work(app)

	ticker := time.NewTicker(time.Minute * time.Duration(app.Config.ReRunInterval))
	defer ticker.Stop()
	for {
		select {
		case t := <-ticker.C:
			app.Logger.Println("Tick at: ", t)
			work(app)
		}
	}

}

func work(app *config.Application) {

	sem := make(chan struct{}, 150)
	var wg sync.WaitGroup

	latest, err := instagram.GetFeed(app)
	if err != nil {
		log.Fatalln(err)
	}

	latest.Next(false)
	for i := len(latest.Items) - 1; i >= 0; i-- {
		item := latest.Items[i]

		wg.Add(1)

		// acquire semaphore
		sem <- struct{}{}

		go func() {
			defer wg.Done()

			itemCode := item.Code

			if store.KeyExists(app, itemCode) {
				app.Logger.Println(itemCode, " already exists. Skipping ...")
				// release semaphore
				<-sem
				return
			}

			ext := ".jpg"
			if item.MediaToString() == "video" {
				ext = ".mp4"
			}

			// Note that this does not download instagram carousel images.
			imgdl, viddl, err := item.Download(app.Config.TempMediaPath, fmt.Sprint(itemCode, ext))
			app.Logger.Println("Download Log.", "Img - ", imgdl, "Vid - ", viddl, "Err - ", err)

			if err == nil {
				if imgdl != "" {
					err = telegram.SendImage(app, "./"+imgdl, item.Caption.Text)
				}
				if viddl != "" {
					err = telegram.SendVideo(app, "./"+viddl, item.Caption.Text)
				}
				// if send didn't work
				if err != nil {
					app.Logger.Println(itemCode, "error sending to telegram", err)
				} else {
					app.Logger.Println(itemCode, "sent to telegram")
					// mark as downloaded and sent
					err = store.SaveKeyVal(app, itemCode, "-")
					if err != nil {
						app.Logger.Println(itemCode, "error saving to db")
					}
				}
			}

			// Delete stuff if you want, but I like to leave it as it is.

			// release semaphore
			<-sem

		}()

	}

	wg.Wait()
}
