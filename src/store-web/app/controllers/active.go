package controllers

import (
	"github.com/revel/revel"
)

// Active .
type Active struct {
	*revel.Controller
}

// Index .
func (c Active) Index() revel.Result {
	return c.Render()
}
