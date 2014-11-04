package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	GorpController
}

func (c App) Index() revel.Result {
	records := map[string]string { "status": "ok" }
	return c.RenderJson(records)
}
