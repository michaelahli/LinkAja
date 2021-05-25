package controllers

import (
	"LinkaAja/packages"
	"net/http"
)

func (c *ctrl) TransferBalance(w http.ResponseWriter, r *http.Request) {
	sender, err := c.uc.GetAccountID(r.URL.Path)
	if err != nil {
		packages.BasicResponse(w, "Error Fetch Sender Account ID from URL", http.StatusBadRequest, err)
		return
	}

	req, message, code, err := c.uc.ProcessTransferRequest(r)
	if err != nil {
		packages.BasicResponse(w, message, code, err)
		return
	}

	message, code, err = c.uc.TransferBalance(sender, req)
	if err != nil {
		packages.BasicResponse(w, message, code, err)
		return
	}

	packages.BasicResponse(w, message, code, nil)
}
