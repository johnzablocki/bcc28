package data

import (
	"encoding/json"

	bolt "github.com/coreos/bbolt"
)

type (
	//BookStore stores books
	BookStore struct {
		store      *Store
		bucketName []byte
	}
)

//Save upserts a book
func (bs *BookStore) Save(b *Book) error {
	return bs.store.Update(func(tx *bolt.Tx) error {
		value, err := json.Marshal(b)
		if err != nil {
			return err
		}
		tx.Bucket(bs.bucketName).Put([]byte(b.ID), value)
		return nil
	})
}

//Remove removes a key
func (bs *BookStore) Remove(b *Book) error {
	return bs.store.Update(func(tx *bolt.Tx) error {
		tx.Bucket(bs.bucketName).Delete([]byte(b.ID))
		return nil
	})
}

//FindByID finds a book by ID
func (bs *BookStore) FindByID(id string, book *Book) error {
	return bs.store.View(func(tx *bolt.Tx) error {
		bookBytes := tx.Bucket(bs.bucketName).Get([]byte(id))
		err := json.Unmarshal(bookBytes, book)
		if err != nil {
			return err
		}
		return nil
	})
}

//FindAll finds all books
func (bs *BookStore) FindAll(books *[]Book) error {
	return bs.store.View(func(tx *bolt.Tx) error {
		tx.Bucket(bs.bucketName).ForEach(func(k, v []byte) error {
			var book Book
			err := json.Unmarshal(v, &book)
			if err != nil {
				return err
			}
			book.ID = string(k)
			*books = append(*books, book)
			return nil
		})
		return nil
	})
}

//Init initializes thing
func (bs *BookStore) Init(config *Config) error {
	db, err := bolt.Open(config.DBName, 0600, nil)
	if err != nil {
		return err
	}
	bs.store.db = db
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("books"))
		if err != nil {
			return err
		}
		return nil
	})
}

//Shutdown closes the db
func (bs *BookStore) Shutdown() error {
	return bs.store.db.Close()
}

//NewBookStore BookStore factory method
func NewBookStore() (*BookStore, error) {
	return &BookStore{
		store: &Store{
			db: nil,
		},
		bucketName: []byte("books"),
	}, nil
}
