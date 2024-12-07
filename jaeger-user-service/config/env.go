package config

import (
	"fmt"
	"jarger-user-service/helper"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var (
	APP_PORT             string
	ROOT_PATH            string
	PSQL_CONNECTION_USER string
	DRIVER_NAME          string

	SERVICE_CLIENT_AUTH_GRPC_ADDRESS string
	GRPC_TIMEOUT                     int
	GRPC_MAX_RECEIVE_SIZE            int

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

	DRIVER_NAME = helper.GetENV("DRIVER_NAME", "")
	PSQL_CONNECTION_USER = helper.GetENV("PSQL_CONNECTION_USER", "")

	SERVICE_CLIENT_AUTH_GRPC_ADDRESS = helper.GetENV("SERVICE_CLIENT_AUTH_GRPC_ADDRESS", "")
	if GRPC_TIMEOUT, err = strconv.Atoi(helper.GetENV("GRPC_TIMEOUT", "")); err != nil {
		panic(fmt.Sprintf("Error while convert type: %s", err.Error()))
	}
	if GRPC_MAX_RECEIVE_SIZE, err = strconv.Atoi(helper.GetENV("GRPC_MAX_RECEIVE_SIZE", "")); err != nil {
		panic(fmt.Sprintf("Error while convert type: %s", err.Error()))
	}
}
