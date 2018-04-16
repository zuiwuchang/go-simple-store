package controllers

import (
	"errors"
	"github.com/revel/revel"
	"store-web/app/utils"
)

func init() {
	revel.InterceptFunc(checkRoot, revel.BEFORE, &Root{})
}
func checkRoot(c *revel.Controller) revel.Result {
	if flasg, ok := c.Session[utils.SessionKeyGroup]; ok &&
		utils.HasFlags(flasg, utils.GroupRootFlag) {
		return nil
	}
	return c.RenderError(errors.New(c.Message("E.PermissionDenied")))
}
