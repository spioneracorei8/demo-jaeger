package main

import (
	"jarger-user-service/config"
	"jarger-user-service/server"
)

func getMainServer() server.Server {
	return server.Server{
		APP_PORT:                         config.APP_PORT,
		ROOT_PATH:                        config.ROOT_PATH,
		GRPC_PORT:                        config.GRPC_PORT,
		DRIVER_NAME:                      config.DRIVER_NAME,
		PSQL_CONNECTION_USER:             config.PSQL_CONNECTION_USER,
		SERVICE_SERVER_AUTH_GRPC_ADDRESS: config.SERVICE_SERVER_AUTH_GRPC_ADDRESS,
		GRPC_TIMEOUT:                     config.GRPC_TIMEOUT,
		GRPC_MAX_RECEIVE_SIZE:            config.GRPC_MAX_RECEIVE_SIZE,
	}
}

func main() {
	server := getMainServer()
	server.Start()
}
