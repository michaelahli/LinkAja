package usecases

import (
	"errors"
	"strconv"
	"strings"
)

func (u *uc) GetAccountID(urlstring string) (int, error) {
	urlPart := strings.Split(urlstring, "/")
	account_id, err := strconv.Atoi(urlPart[2])
	if err != nil {
		return 0, errors.New("Failed to Get Params")
	}
	return account_id, nil
}
