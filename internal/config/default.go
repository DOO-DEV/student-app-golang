package config

import (
	"student-app/internal/db"
	"time"
)

func Default() Config {
	return Config{
		Debug:  true,
		Secret: "secret",
		Database: db.Config{
			URL:               "mongodb://127.0.0.1:27017",
			Name:              "students",
			ConnectionTimeout: time.Second,
		},
		Admin: Admin{
			Username: "admin",
			Password: "admin",
			Name:     "hossein",
		},
	}
}
