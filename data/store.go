package data

import (
	"github.com/coreos/bbolt"
)

type (
	//Store represents the model store base
	Store struct {
		db *bolt.DB
	}
)

//Update mutates the database
func (store *Store) Update(fn func(tx *bolt.Tx) error) error {
	return store.db.Update(fn)
}

//View reads the database
func (store *Store) View(fn func(tx *bolt.Tx) error) error {
	return store.db.View(fn)
}
