package database

import (
	"fmt"
	"log"
	"time"

	"github.com/dusk-chancellor/mego-favorite/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)


func ConnectToDB(cfg *config.Config) (*sqlx.DB, error) {
	var (
		db *sqlx.DB
		err error
		maxAttempts = 10
		databaseURL = fmt.Sprintf(
			"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
			cfg.DBUser,
			cfg.DBPassword,
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBName,
		)
	)

	for i := 0; i < maxAttempts; i++ {
		db, err = sqlx.Connect("postgres", databaseURL)
		if err == nil {
			log.Printf("Connected to database")
			return db, nil
		}
		log.Printf("Failed to connect to database, retrying...")
		time.Sleep(5*time.Second)
	}
	log.Printf("Failed to connect to database after %d attempts", maxAttempts)
	return nil, err
}