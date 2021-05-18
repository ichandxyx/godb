package store

import (
	"context"

	"github.com/ichandxyx/godb/ent"
)

func (s *Store) GetUsers(ctx context.Context) ([]*ent.User, error) {
	return s.db.User.Query().All(ctx)
}
func (s *Store) AddUsers(ctx context.Context, name string, age int)(*ent.User, error){
	return s.db.User.Create().
	SetAge(age).
	SetName(name).
	Save(ctx)
}
