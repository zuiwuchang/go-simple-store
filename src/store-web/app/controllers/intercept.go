package controllers

import (
	"errors"
	"github.com/revel/revel"
	"store-web/app/utils"
)

func init() {
	revel.InterceptFunc(checkRoot, revel.BEFORE, &Root{})
	revel.InterceptFunc(checkActive, revel.BEFORE, &Active{})
}
func checkRoot(c *revel.Controller) revel.Result {
	if flasg, ok := c.Session[utils.SessionKeyGroup]; ok &&
		utils.HasFlags(flasg, utils.GroupRootFlag) {
		if active, ok := c.Session[utils.SessionKeyActive]; ok && active == "1" {
			return nil
		}
	}
	return c.RenderError(errors.New(c.Message("E.PermissionDenied")))
}
func checkActive(c *revel.Controller) revel.Result {
	if val, ok := c.Session[utils.SessionKeyActive]; ok && val == "0" {
		return nil
	}
	return c.Redirect("/")
}
