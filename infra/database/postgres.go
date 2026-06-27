package database

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gundz204/pglib/odm/base"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
)

func Connect() (*sqlx.DB, error) {
	var db *sqlx.DB
	var err error

	cfg := *base.InitConfig()
	vip := cfg.GetConfig()

	mode := vip.GetString("database.postgres.mode")

	switch mode {
	case "remote":
		serviceURI := vip.GetString("database.postgres.url")

		conn, err := url.Parse(serviceURI)
		if err != nil {
			return nil, err
		}

		conn.RawQuery = "sslmode=verify-ca&sslrootcert=ca.pem"

		db, err = sqlx.Open("postgres", conn.String())
		if err != nil {
			return nil, err
		}

	default:
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			vip.GetString("database.postgres.host"),
			vip.GetString("database.postgres.port"),
			vip.GetString("database.postgres.user"),
			vip.GetString("database.postgres.password"),
			vip.GetString("database.postgres.db_name"),
			vip.GetString("database.postgres.db_sslmode"),
		)

		db, err = sqlx.Open("pgx", dsn)
		if err != nil {
			return nil, err
		}
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("database connected")

	return db, nil
}
