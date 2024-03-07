package main

import (
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

type Env struct {
	ServerPort string
}

func LoadEnv() Env {
	env, err := filepath.Abs("../.env")

	if err != nil {
		panic(err)
	}

	if err = godotenv.Load(env); err != nil {
		panic(err)
	}

	return Env{
		ServerPort: os.Getenv("SERVER_PORT"),
	}
}
