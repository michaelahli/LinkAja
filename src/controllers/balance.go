package controllers

import (
	"LinkaAja/packages"
	"net/http"
)

func (c *ctrl) GetBalance(w http.ResponseWriter, r *http.Request) {
	account_number, err := c.uc.GetAccountID(r.URL.Path)
	if err != nil {
		packages.BasicResponse(w, "Error Fetch Account ID from URL", http.StatusBadRequest, err)
		return
	}

	res, err := c.uc.GetBalanceInfo(account_number)
	if err != nil {
		packages.BasicResponse(w, "Failed to Get Balance", http.StatusNotFound, err)
		return
	}

	packages.BasicResponse(w, "Success", http.StatusOK, res)
	return
}
