package main

import (
	"jaeger-auth-service/config"
	"jaeger-auth-service/server"
)

func getMainServer() server.Server {
	return server.Server{
		ROOT_PATH:                        config.ROOT_PATH,
		APP_PORT:                         config.APP_PORT,
		GRPC_PORT:                        config.GRPC_PORT,
		GRPC_TIMEOUT:                     config.GRPC_TIMEOUT,
		GRPC_MAX_RECEIVE_SIZE:            config.GRPC_MAX_RECEIVE_SIZE,
		SERVICE_SERVER_USER_GRPC_ADDRESS: config.SERVICE_SERVER_USER_GRPC_ADDRESS,
		DRIVER_NAME:                      config.DRIVER_NAME,
		PSQL_CONNECTION_AUTH:             config.PSQL_CONNECTION_AUTH,
	}
}

func main() {
	server := getMainServer()
	server.Start()
}
