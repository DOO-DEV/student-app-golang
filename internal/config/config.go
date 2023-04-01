package config

import "student-app/internal/db"

type Config struct {
	Debug    bool      `koanf:"debug"`
	Secret   string    `koanf:"secret"`
	Database db.Config `koanf:"database"`
	Admin    Admin     `koanf:"admin"`
}

type Admin struct {
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	Name     string `koanf:"name"`
}
