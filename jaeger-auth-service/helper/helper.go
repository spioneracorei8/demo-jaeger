package helper

import (
	"fmt"
	"jaeger-auth-service/constant"
	"os"
	"time"
)

func NewTimestampFromTime(t time.Time) time.Time {
	location := time.FixedZone("UTC+7", 7*60*60)
	now, err := time.Parse(constant.TIMESTAMP_LAYOUT, t.UTC().Format(constant.TIMESTAMP_LAYOUT))
	if err != nil {
		panic(fmt.Sprintf("Error while parse time : %s", err.Error()))
	}
	now = now.In(location)
	return now
}

func GetENV(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
