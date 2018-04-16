package controllers

import (
	"github.com/revel/revel"
	kEmail "github.com/zuiwuchang/king-go/net/email"
	"store-web/app/ajax"
	"store-web/app/configure"
	"store-web/app/db/data"
	"store-web/app/db/manipulator"
	"strings"
)

// Root 僅root 可以訪問的 管理頁
type Root struct {
	*revel.Controller
}

// System 系統管理
func (c Root) System() revel.Result {
	version := configure.Get().Version
	var mSys manipulator.SystemInfo
	systemInfo, e := mSys.Get()
	if e != nil {
		return c.RenderError(e)
	}
	return c.Render(version, systemInfo)
}

// AjaxSaveRegister 修改 註冊模式
func (c Root) AjaxSaveRegister(registerMode int) revel.Result {
	var result ajax.Result

	var mSys manipulator.SystemInfo
	if e := mSys.SaveRegister(registerMode); e != nil {
		result.Code = ajax.Error
		result.Emsg = e.Error()
	}

	return c.RenderJSON(&result)
}

// AjaxSaveSMTP 修改 smtp 配置
func (c Root) AjaxSaveSMTP(smtp, email, pwd string) revel.Result {
	var result ajax.Result

	var mSys manipulator.SystemInfo
	if e := mSys.SaveSMTP(smtp, email, pwd); e != nil {
		result.Code = ajax.Error
		result.Emsg = e.Error()
	}

	return c.RenderJSON(&result)
}

// AjaxTestSMTP 測試 smtp 配置
func (c Root) AjaxTestSMTP(smtp, email, pwd string) revel.Result {
	var result ajax.Result
	client, e := kEmail.NewSMTPSSLClient(smtp, email, pwd)
	if e == nil {
		client.Quit()
	} else {
		result.Code = ajax.Error
		result.Emsg = e.Error()
	}
	return c.RenderJSON(&result)
}

// AjaxSaveActive 存儲 激活郵件 配置
func (c Root) AjaxSaveActive(title, text string) revel.Result {
	var result ajax.Result

	var mSys manipulator.SystemInfo
	if e := mSys.SaveActive(title, text); e != nil {
		result.Code = ajax.Error
		result.Emsg = e.Error()
	}

	return c.RenderJSON(&result)
}

// AjaxTestActive .
func (c Root) AjaxTestActive(text string) revel.Result {
	var result ajax.Result
	text = strings.TrimSpace(text)

	if text == "" {
		result.Code = ajax.Error
		result.Emsg = c.Message("rSystem.e.active text empty")
		return c.RenderJSON(&result)
	}

	str, e := data.GetActiveEmail(
		&data.ActiveContext{
			Host:  c.Request.Host,
			Email: "testName@xxx.xxx",
			ID:    0,
		},
		text,
	)
	if e != nil {
		result.Code = ajax.Error
		result.Emsg = e.Error()

		return c.RenderJSON(&result)
	}
	result.Str = str
	return c.RenderJSON(&result)
}
