package handler

import (
	"encoding/json"
	"jaeger-auth-service/models"
)

func unmarshal(s string) (*models.Account, error) {
	var (
		account = new(models.Account)
		err     error
	)
	if err = json.Unmarshal([]byte(s), &account); err != nil {
		return nil, err
	}
	return account, nil

}
