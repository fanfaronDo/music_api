package app

import (
	"database/sql"
	"github.com/fanfaronDo/music_api/pkg/config"
)

type DataBase interface {
	Connect(host string, cfg *config.Config) (*sql.DB, error)
}

func Run(cnf *config.Config) error {

	return nil
}
