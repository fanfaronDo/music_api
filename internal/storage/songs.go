package storage

import (
	"database/sql"
	"github.com/fanfaronDo/music_api/internal/entry"
)

type SongsStorage struct {
	db *sql.DB
}

func NewSongsStorage(db *sql.DB) *SongsStorage {
	return &SongsStorage{db: db}
}

func (s *SongsStorage) CreateSong(song *entry.Song) (*entry.Song, error) {

}
