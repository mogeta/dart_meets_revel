package controllers

import (
	"github.com/robfig/revel"
)

type Sandbox struct {
	*revel.Controller
}

func (c Sandbox) D3() revel.Result {
	return c.Render()
}
