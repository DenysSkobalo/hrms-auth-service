package config

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
)

var once sync.Once

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
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
}

func ConnectDB() (*sql.DB, error) {
	var db *sql.DB
	var err error

	once.Do(func() {
		config := loadDatabaseConfig()
		connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s", config.Username, config.Password, config.Host, config.Port, config.DBName, config.SSLMode)

		var err error
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			err = fmt.Errorf("failed to connect to database: %v", err)
			return
		}

		err = db.Ping()
		if err != nil {
			err = fmt.Errorf("failed to ping database: %v", err)
			db.Close()
			db = nil
			return
		}

	})

	return db, err
}
