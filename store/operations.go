package store

import (
	"errors"

	"github.com/IceWreck/InstaTG/config"
	"github.com/boltdb/bolt"
)

// KeyExists checks if your key exists in db
func KeyExists(app *config.Application, key string) bool {

	exists := false

	app.BoltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("instatg"))
		if b == nil {
			app.Logger.Println("Bucket instatg does not exist")
			return nil
		}
		v := b.Get([]byte(key))
		if v != nil {
			exists = true
		}
		return nil
	})

	return exists
}

// SaveKeyVal - save your key, value to db
func SaveKeyVal(app *config.Application, key string, value string) error {

	return app.BoltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("instatg"))
		if b == nil {
			app.Logger.Println("Bucket instatg does not exist")
			return errors.New("bucket instatg does not exist")
		}
		return b.Put([]byte(key), []byte(value))
	})
}
