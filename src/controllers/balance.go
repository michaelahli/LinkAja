package controllers

import (
	"LinkaAja/packages"
	"LinkaAja/src/models"
	"net/http"
)

func (c *ctrl) GetBalance(w http.ResponseWriter, r *http.Request) {
	account_number, err := c.uc.GetAccountID(r.URL.Path)
	if err != nil {
		packages.BasicResponse(w, "Error Fetch Account ID from URL", http.StatusBadRequest, err)
		return
	}

	res := models.GetBalance{}
	res.AccountNumber = account_number

	packages.BasicResponse(w, "Success", http.StatusOK, res)
	return
}
