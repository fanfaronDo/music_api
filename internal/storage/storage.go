package storage

import "database/sql"

type Songs interface {
}

type Storage struct {
	Songs
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		NewSongsStorage(db),
	}
}
