package controllers

import (
	"github.com/revel/revel"
	"github.com/miry/revel_mysql/app/models"
)

type Accounts struct {
	App
}

func (c Accounts) Index() revel.Result {
	var accounts []*models.Account
	_, err := c.Txn.Select(&accounts, "SELECT id, name, created_at FROM accounts")

	if err != nil {
		return c.RenderJson(map[string]string{"status": "error", "message": err.Error()})
	}

	return c.RenderJson(accounts)
}

func (c Accounts) Show(id int64) revel.Result {
	revel.INFO.Print(id)

	account := new(models.Account)
	err := c.Txn.SelectOne(&account,
		"SELECT id, name, created_at FROM accounts WHERE id = ?", id)

	if err != nil {
		revel.INFO.Print(err)
		return c.RenderText("Error. Item probably doesn't exist.")
	}

	return c.RenderJson(account)
}
