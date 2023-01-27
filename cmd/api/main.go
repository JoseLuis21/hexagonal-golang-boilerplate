package main

import (
	"hexagonal-go/cmd/api/bootstrap"
	"hexagonal-go/config"
	kit "hexagonal-go/kit/mysql"
	"log"
)

func main() {
	
	envConfig := config.GetConfig()

	db, err := kit.MysqlCreateConnection(envConfig.DbUser,envConfig.DbPass,envConfig.DbHost,
		envConfig.DbPort,envConfig.DbName)
	if err != nil {
		log.Fatal(err)
	}
	
	if err := bootstrap.RunServer(db, envConfig.AppName, envConfig.Port,
		envConfig.DbTimeout,envConfig.ShutdownTimeout); err != nil {
		log.Fatal(err)
	}

}
