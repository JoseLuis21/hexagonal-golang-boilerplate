package bootstrap

import (
	"database/sql"
	"hexagonal-go/internal/application"
	"hexagonal-go/internal/infraestructure"
	"hexagonal-go/internal/server"
	"time"
)

func RunServer(db *sql.DB, appName string, port uint32, dbTimeout time.Duration, shutdownTimeout time.Duration) error {
	userRepository := infraestructure.NewUserRepository(db, dbTimeout)
	createUserService := application.NewUserCreateService(userRepository)
	srv := server.New(appName, port, shutdownTimeout, createUserService)
	return srv.Run()
}
