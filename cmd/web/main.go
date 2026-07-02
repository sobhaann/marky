package main

import "github.com/sobhaann/markdown-notetaking-api/internal/env"

type application struct {
	config *env.Config
}

func NewApplication() *application {
	return &application{
		config: env.NewConfig(),
	}
}

func main() {
	NewApplication()
}
