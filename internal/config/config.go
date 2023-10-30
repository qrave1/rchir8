package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

type Config struct {
	Address  string
	Port     string
	Cookie   string
	Hashkey  []byte
	BlockKey []byte
}

func NewConfig() *Config {
	return &Config{
		Address:  os.Getenv("ADDRESS"),
		Port:     os.Getenv("PORT"),
		Cookie:   os.Getenv("COOKIE"),
		Hashkey:  []byte(os.Getenv("HASHKEY")),
		BlockKey: []byte(os.Getenv("BLOCK_KEY")),
	}
}
