package configs

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"sync"
)

var (
	db   *sql.DB
	once sync.Once
)

type DBConfig struct {
	Host     string
	Username string
	Password string
	Port     string
	DBName   string
	SSLMode  string
}

func loadDatabaseConfig() DBConfig {
	return DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
}

func ConnectDB() *sql.DB {
	once.Do(func() {
		config := loadDatabaseConfig()
		connStr := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s",
			config.Host, config.Username, config.Password, config.Port, config.DBName, config.SSLMode)

		var err error
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}

		err = db.Ping()
		if err != nil {
			log.Fatal("Error connecting to database: ", err)
		}
	})

	return db
}
