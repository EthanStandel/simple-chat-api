package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DatabaseUrl interface {
	url() string
}

type Database struct {
	username, password, host, port, db, sslMode string
}

func (db Database) url() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		db.username,
		db.password,
		db.host,
		db.port,
		db.db,
	)
}

var database = Database{
	username: os.Getenv("SCA_DB_USERNAME"),
	password: os.Getenv("SCA_DB_PASSWORD"),
	host:     os.Getenv("SCA_DB_HOST"),
	port:     os.Getenv("SCA_DB_PORT"),
	db:       os.Getenv("SCA_DB_DB"),
	sslMode:  os.Getenv("SCA_DB_SSLMODE"),
}

var pool *pgxpool.Pool

func GetConn(ctx context.Context) *pgx.Conn {
	if pool == nil {
		pool, _ = pgxpool.Connect(ctx, database.url())
	}

	conn, _ := pool.Acquire(ctx)
	return conn.Conn()
}

func CloseConnection() {
	if pool != nil {
		pool.Close()
	}
}
