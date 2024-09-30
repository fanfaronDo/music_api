package app

import (
	"context"
	"fmt"
	"github.com/fanfaronDo/music_api/internal/storage"
	"github.com/fanfaronDo/music_api/pkg/config"
)

func Run(cnf *config.Config) error {
	ctx := context.Background()
	db, err := storage.NewPostgres(ctx, cnf.Postgres.Host, cnf)

	if err != nil {
		return err
	}

	fmt.Println(db)

	return nil
}
