package main

import (
	"jaeger-auth-service/config"
	"jaeger-auth-service/server"
)

func getMainServer() server.Server {
	return server.Server{
		APP_PORT:             config.APP_PORT,
		ROOT_PATH:            config.ROOT_PATH,
		DRIVER_NAME:          config.DRIVER_NAME,
		PSQL_CONNECTION_AUTH: config.PSQL_CONNECTION_AUTH,
	}
}

func main() {
	server := getMainServer()
	server.Start()
}
