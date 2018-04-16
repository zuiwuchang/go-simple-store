package controllers

import (
	"github.com/revel/revel"
	"store-web/app/utils"
)

// 註冊 模板
func init() {
	revel.TemplateFuncs["HasFlags"] = func(iFlags, iFlag interface{}) bool {
		flags, ok := iFlags.(string)
		if !ok {
			return false
		}
		flag, ok := iFlag.(string)
		if !ok {
			return false
		}
		return utils.HasFlags(flags, flag)
	}
}
