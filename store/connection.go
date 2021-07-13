package store

import (
	"time"

	"github.com/IceWreck/InstaTG/config"
	"github.com/boltdb/bolt"
)

// BoldDB is an embedded key value store

// NewConnection returns a new boltdb connection
func NewConnection(app *config.Application) *bolt.DB {
	db, err := bolt.Open(app.Config.DatabasePath, 0600, nil)
	if err != nil {
		app.Logger.Fatalln("Failed to open db ", err)
	}
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("instatg"))
		if err != nil {
			return err
		}
		return b.Put([]byte("last_opened"), []byte(time.Now().String()))
	})

	// remember to close this wherever you call store.NewConnection()
	return db
}
