package storage

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/fanfaronDo/music_api/pkg/config"
	"strconv"
	"time"
)

func NewPostgres(ctx context.Context, host string, cfg *config.Config) (*sql.DB, error) {
	if cfg == nil {
		return nil, ErrConfigNull
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer func() {
		fmt.Println("Canceled with context.Context")
		cancel()
	}()

	portINT, err := strconv.Atoi(cfg.Postgres.Port)
	if err != nil {
		return nil, err
	}

	if portINT <= 0 {
		return nil, ErrConfigPortLessZero
	}

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		host, portINT, cfg.Postgres.User, cfg.Postgres.Database, cfg.Postgres.Password, cfg.Postgres.SSLMode))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
