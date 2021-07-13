package instagram

import (
	"fmt"
	"sync"

	"github.com/IceWreck/InstaTG/config"
	"github.com/ahmdrz/goinsta/v2"
)

type InstagramPost struct {
	Code      string
	Link      string
	Caption   string
	MediaType string
	Time      int64
}

// GetFeed returns a goinsta.FeedMedia struct
// call latest.Next() to fill it up.
func GetFeed(app *config.Application) (*goinsta.FeedMedia, error) {
	user, err := app.InstaBot.Profiles.ByName(app.Config.InstagramChannel)
	if err != nil {
		app.Logger.Println("Unable to fetch user. Error: ", err)
		return nil, err
	}
	return user.Feed(), nil
}

// Downloads the latest Posts to the app.Config.TempMediaPath folder.
// This function will probably be unused in the final version and
// functionality will be moved to cmd after combining with tg.
// Am leaving it here as an example.
func DownloadPosts(app *config.Application) error {
	user, err := app.InstaBot.Profiles.ByName(app.Config.InstagramChannel)
	if err != nil {
		app.Logger.Println("Unable to fetch user. Error: ", err)
		return err
	}

	items := []InstagramPost{}

	latest := user.Feed()
	if err != nil {
		app.Logger.Println("Unable to fetch feed. Error: ", err)
		return err
	}

	wg := new(sync.WaitGroup)
	latest.Next()

	for i := len(latest.Items) - 1; i >= 0; i-- {
		item := latest.Items[i]

		if len(item.Images.Versions) > 0 {
			post := InstagramPost{
				Code:      item.Code,
				Caption:   item.Caption.Text,
				MediaType: item.MediaToString(),
				Time:      item.DeviceTimestamp,
			}
			wg.Add(1)
			go func() {
				defer wg.Done()
				imgdl, viddl, errdl := item.Download(app.Config.TempMediaPath, fmt.Sprint(item.Code))
				app.Logger.Println("Downloaded ", imgdl, viddl, errdl)
			}()
			items = append(items, post)
		}

	}

	if err := latest.Error(); err != nil {

		return fmt.Errorf("unable to retrieve user feed %v", err)
	}

	// sort.Slice(items, func(i, j int) bool { return items[i].Time < items[j].Time })

	fmt.Println(items)
	wg.Wait()
	return nil
}
