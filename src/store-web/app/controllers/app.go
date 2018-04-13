package controllers

import (
	"github.com/revel/revel"
	"store-web/app/ajax"
	"store-web/app/configure"
	"store-web/app/db/manipulator"
)

// App .
type App struct {
	*revel.Controller
}

// Index 首頁
func (c App) Index() revel.Result {
	version := configure.Get().Version
	return c.Render(version)
}

// About 關於
func (c App) About() revel.Result {
	return c.Render()
}

// License 許可協議
func (c App) License() revel.Result {
	return c.Render()
}

// Register 註冊
func (c App) Register() revel.Result {
	var mSys manipulator.SystemInfo
	systemInfo, e := mSys.Get()
	if e != nil {
		return c.RenderError(e)
	}
	return c.Render(systemInfo)
}

// AjaxIsEmailExists 驗證 email 是否 存在
func (c App) AjaxIsEmailExists(email string) revel.Result {
	var result ajax.Result
	var mUser manipulator.User
	if yes, e := mUser.IsEmailExists(email); e != nil {
		result.Code = ajax.Error
		result.Emsg = e.Error()
	} else if yes {
		result.Value = 1
	}
	return c.RenderJSON(&result)
}

// AjaxRegister 註冊 用戶 成功 返回 用戶 id
func (c App) AjaxRegister(email, pwd, code string) revel.Result {
	var result ajax.Result
	var mUser manipulator.User
	if user, e := mUser.Register(email, pwd, code); e != nil {
		result.Code = ajax.Error
		result.Emsg = e.Error()
	} else {
		result.Value = user.ID
		// 設置 登入 session
	}
	return c.RenderJSON(&result)
}

// Login 登入
func (c App) Login() revel.Result {
	return c.Render()
}
