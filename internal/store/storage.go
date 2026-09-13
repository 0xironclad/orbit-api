package store

import (
	"context"
	"database/sql"
	
)

type Storage struct {
	Posts interface {
		Create(context context.Context) error
	}
	Users interface {
		Create(context context.Context) error
		
	}	
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{db: db},
		Users: &UserStore{db: db},
	}
}