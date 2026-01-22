package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error cargando .env: %v\n", err)
	}
}
