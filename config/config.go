package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type EnvConfigApp struct {
	AppName string
	Port   uint32
	DbUser string
	DbPass string
	DbHost string
	DbPort string
	DbName string
	DbTimeout time.Duration
	ShutdownTimeout time.Duration
}

func generateEnvirommentConfig() *EnvConfigApp {
	godotenv.Load("../../config/.env")
	port, _ := strconv.ParseUint(os.Getenv("PORT"),10,64)

	fmt.Println(os.Getenv("PORT"))
	dbtimeout,_ := time.ParseDuration(os.Getenv("DB_TIMEOUT"))
	shutdownTimeout, _ := time.ParseDuration(os.Getenv("SHUTDOWN_TIMEOUT"))
	return &EnvConfigApp{
		AppName: os.Getenv("APP_NAME"),
		Port:  uint32(port),
		DbUser: os.Getenv("DB_USER"),
		DbPass: os.Getenv("DB_PASS"),
		DbHost: os.Getenv("DB_HOST"),
		DbPort: os.Getenv("DB_PORT"),
		DbName: os.Getenv("DB_NAME"),
		DbTimeout: dbtimeout,
		ShutdownTimeout: shutdownTimeout,
	}
}

func GetConfig() *EnvConfigApp {
	return generateEnvirommentConfig()
}