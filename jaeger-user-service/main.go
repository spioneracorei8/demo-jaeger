package main

import (
	"jarger-user-service/config"
	"jarger-user-service/server"
)

func getMainServer() server.Server {
	return server.Server{
		APP_PORT:             config.APP_PORT,
		ROOT_PATH:            config.ROOT_PATH,
		DRIVER_NAME:          config.DRIVER_NAME,
		PSQL_CONNECTION_USER: config.PSQL_CONNECTION_USER,
	}
}

func main() {
	server := getMainServer()
	server.Start()
}
