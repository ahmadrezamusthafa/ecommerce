package main

import (
	"github.com/ahmadrezamusthafa/ecommerce/cmd"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/logger"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.Errorf("Error loading .env file | %v", err)
	}
	logger.Init()
}

func main() {
	cmd.Execute()
}
