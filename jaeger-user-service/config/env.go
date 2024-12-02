package config

import (
	"fmt"
	"jarger-user-service/helper"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
)

var (
	APP_PORT             string
	ROOT_PATH            string
	PSQL_CONNECTION_USER string
	DRIVER_NAME          string

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
}
