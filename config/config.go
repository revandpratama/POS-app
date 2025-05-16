package config

import "os"

type Config struct {
	PORT string

	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string

	JWT_SECRET string
}

var ENV *Config

func InitConfig() {
	ENV = &Config{
		PORT: os.Getenv("PORT"),

		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),

		JWT_SECRET: os.Getenv("JWT_SECRET"),
	}
}
