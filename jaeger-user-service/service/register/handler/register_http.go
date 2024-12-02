package handler

import (
	"encoding/json"
	"jarger-user-service/models"
)

func unmarshal(s string) (*models.User, error) {
	var (
		user = new(models.User)
		err     error
	)
	if err = json.Unmarshal([]byte(s), &user); err != nil {
		return nil, err
	}
	return user, nil

}
