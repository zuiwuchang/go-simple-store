package controllers

import (
	"github.com/revel/revel"
	"store-web/app/configure"
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
	return c.Render()
}

// Login 登入
func (c App) Login() revel.Result {
	return c.Render()
}