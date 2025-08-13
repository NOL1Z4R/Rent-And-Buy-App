package config

import (
	"bytes"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Connection     string
	JwtSecret      string
	JwtExpireHours int
}

// In case of some IDEs having errors during fetching the Env File.
func DefaultToEnvFile() {
	envPath := "./.env"
	content, err := os.ReadFile(envPath)
	if err != nil {
		fmt.Println("Error during reading file.", err)
		return
	}

	bom := []byte{0xEF, 0xBB, 0xBF}
	content = bytes.TrimPrefix(content, bom)
	err = os.WriteFile(envPath, content, 0644)
	if err != nil {
		fmt.Println("Error during writing to file.", err)
		return
	}
}

func LoadConfig() *Config {
	DefaultToEnvFile()

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	hours, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE_HOURS"))

	config := &Config{
		Connection:     os.Getenv("CONNECTION"),
		JwtSecret:      os.Getenv("JWT_SECRET"),
		JwtExpireHours: hours,
	}

	return config
}

func GetDsn(config *Config) string {
	return fmt.Sprintf(config.Connection)
}
