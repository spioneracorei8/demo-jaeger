package config

import (
	"fmt"
	"jaeger-auth-service/helper"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var (
	ROOT_PATH                        string
	APP_PORT                         string
	GRPC_PORT                        string
	GRPC_TIMEOUT                     int
	GRPC_MAX_RECEIVE_SIZE            int
	SERVICE_SERVER_USER_GRPC_ADDRESS string
	PSQL_CONNECTION_AUTH             string
	DRIVER_NAME                      string

	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func init() {
	var (
		index int
		err   error
	)
	if err = godotenv.Load(); err != nil {
		panic(fmt.Sprintf("Error while loading .env: %s", err.Error()))
	}
	index = strings.LastIndex(basepath, "/config")
	if index != -1 {
		ROOT_PATH = strings.Replace(basepath, "/config", "", index)
	}
	APP_PORT = helper.GetENV("APP_PORT", "")
	GRPC_PORT = helper.GetENV("GRPC_PORT", "")
	if GRPC_TIMEOUT, err = strconv.Atoi(helper.GetENV("GRPC_TIMEOUT", "")); err != nil {
		panic(fmt.Sprintf("Error while converting type:%s", err.Error()))
	}
	if GRPC_MAX_RECEIVE_SIZE, err = strconv.Atoi(helper.GetENV("GRPC_MAX_RECEIVE_SIZE", "")); err != nil {
		panic(fmt.Sprintf("Error while converting type:%s", err.Error()))
	}
	SERVICE_SERVER_USER_GRPC_ADDRESS = helper.GetENV("SERVICE_SERVER_USER_GRPC_ADDRESS", "")
	DRIVER_NAME = helper.GetENV("DRIVER_NAME", "")
	PSQL_CONNECTION_AUTH = helper.GetENV("PSQL_CONNECTION_AUTH", "")

}
