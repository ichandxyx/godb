package store

import (
	"github.com/ichandxyx/godb/ent"
)

type Store struct {
	db *ent.Client
}

func New(client *ent.Client) *Store {
	return &Store{
		db: client,
	}
}
