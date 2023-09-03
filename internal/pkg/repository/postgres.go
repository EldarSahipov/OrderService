package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	NameDB   string
}

func NewPostgresDB(conf Config) (*pgxpool.Pool, error) {
	dbUrl := "postgres://" + conf.Username + ":" + conf.Password + "@" + conf.Host + ":" + conf.Port + "/" + conf.NameDB
	db, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}
	return db, nil
}
