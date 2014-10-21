package controllers

import "github.com/revel/revel"
import (
	"bitcup/app/cron"
	"fmt"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	for _, adr := range cron.BtcAddress.Address {
		c.RenderArgs[fmt.Sprintf("%s%s", "btc", adr.Address)] = fmt.Sprintf("%0.3f", adr.Final_balance/100000000.0)
	}
	return c.Render()
}

func (c App) Help() revel.Result {
	return c.Render()
}
func (c App) About() revel.Result {
	return c.Render()
}
