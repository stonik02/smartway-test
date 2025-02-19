package config

import "os"

type Config struct {
	Database DB  `yaml:"database"`
	App      App `yaml:"app"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func (dbConfig *DB) Fill() {
	dbConfig.Host = os.Getenv("PG_HOST")
	if dbConfig.Host == "" {
		dbConfig.Host = "localhost"
	}
	dbConfig.Port = os.Getenv("PG_PORT")
	if dbConfig.Port == "" {
		dbConfig.Port = "5432"
	}
	dbConfig.Username = os.Getenv("PG_USER")
	if dbConfig.Username == "" {
		dbConfig.Username = "postgres"
	}
	dbConfig.Password = os.Getenv("PG_PASSWORD")
	if dbConfig.Password == "" {
		dbConfig.Password = "root"
	}
	dbConfig.Database = os.Getenv("PG_DB_NAME")
	if dbConfig.Database == "" {
		dbConfig.Database = "postgres"
	}
}

type App struct {
	Port string `yaml:"port"`
}

func (appConfig *App) Fill() {
	appConfig.Port = os.Getenv("APP_PORT")
	if appConfig.Port == "" {
		appConfig.Port = "8080"
	}
}

func New() *Config {
	cfg := &Config{}
	cfg.Database.Fill()
	cfg.App.Fill()
	return cfg
}
