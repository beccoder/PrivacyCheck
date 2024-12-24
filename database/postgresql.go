package database

import (
	"errors"
	"privacy-check/configs/pg"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var instance *InstanceHolder

type InstanceHolder struct {
	*sqlx.DB
}

func (i *InstanceHolder) clone() *InstanceHolder {
	return &InstanceHolder{DB: i.DB}
}

// TODO: add max connections | settings
// db.SetConnMaxIdleTime(n)
// db.SetMaxOpenConns(n)
// db.SetMaxIdleConns(n)
// db.SetConnMaxLifetime(n)
func connect(cfg pg.Config) (*InstanceHolder, error) {
	// singleton pattern
	if instance != nil {
		return instance.clone(), nil
	}

	db, err := sqlx.Connect("pgx", cfg.Build())
	if err != nil {
		return nil, err
	}

	instance = &InstanceHolder{db}

	return instance.clone(), nil
}

// max try 20.
// max timeout 10.
// reconnection time cannot be more than 10 seconds one step
// the number of attempts can be less than or equal to 20
// connection time out
func RetryConnect(cfg pg.Config, tryCount uint, timeout uint64) (*InstanceHolder, error) {
	if timeout > 10 {
		return nil, errors.New("reconnection time cannot be more than 10 seconds one step")
	}

	if tryCount > 20 {
		return nil, errors.New("the number of attempts can be less than or equal to 20")
	}

	timeoutSecond := time.Second * time.Duration(timeout)

	for i := uint(1); i <= tryCount; i++ {
		db, err := connect(cfg)

		if err == nil {
			return db, nil
		}

		time.Sleep(timeoutSecond)
	}

	return nil, errors.New("connection time out")
}

// value or error
func Connect(cfg pg.Config) (*sqlx.DB, error) {
	db, err := connect(cfg)
	if err != nil {
		return nil, err
	}

	return db.DB, err
}

// value or panic
func MustConnect(cfg pg.Config) *sqlx.DB {
	db, err := connect(cfg)
	if err != nil {
		panic(err)
	}

	return db.DB
}
