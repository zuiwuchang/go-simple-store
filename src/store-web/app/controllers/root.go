package controllers

import (
	"github.com/revel/revel"
	kEmail "github.com/zuiwuchang/king-go/net/email"
	"store-web/app/ajax"
	"store-web/app/configure"
	"store-web/app/db/data"
	"store-web/app/db/manipulator"
	"strings"
	"time"
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

	str, e := data.GetActiveCode(0, time.Now().Unix())
	if e != nil {
		result.Code = ajax.Error
		result.Emsg = e.Error()

		return c.RenderJSON(&result)
	}
	str, e = data.GetActiveEmail(
		&data.ActiveContext{
			Host:  c.Request.Host,
			Email: "testName@xxx.xxx",
			ID:    0,
			Code:  str,
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

// Code 邀請碼 管理
func (c Root) Code() revel.Result {
	return c.Render()
}

// AjaxCode 創建 邀請碼
func (c Root) AjaxCode() revel.Result {
	var result ajax.Result

	var mCode manipulator.InviteCode
	if code, e := mCode.Code(); e == nil {
		result.Str = code
	} else {
		result.Code = ajax.Error
		result.Emsg = e.Error()
	}
	return c.RenderJSON(&result)
}

// AjaxFindCode .
func (c Root) AjaxFindCode(rows, page, pages int64, code string) revel.Result {
	var result ajax.ResultFindCode
	//每頁顯示 數量
	if rows < 10 {
		rows = 10
	} else if rows > 50 {
		rows = 50
	}

	//當前頁 從1 開始 計數
	if page < 1 {
		page = 1
	}

	//不知道 總頁數
	var mCode manipulator.InviteCode
	if pages < 1 {
		// 查詢 總頁數
		n, e := mCode.Count(code)
		if e != nil {
			result.Code = ajax.Error
			result.Emsg = e.Error()
			return c.RenderJSON(&result)
		}
		if n < 1 {
			pages = 0
		} else {
			pages = (n + rows - 1) / rows
		}
	}
	result.Pages = pages
	if pages < 1 {
		// 沒有數據 直接 返回
		return c.RenderJSON(&result)
	}

	//查詢 數據
	start := (page - 1) * rows
	datas, e := mCode.Find(start, rows, code)
	if e != nil {
		result.Code = ajax.Error
		result.Emsg = e.Error()
		return c.RenderJSON(&result)
	}

	result.Data = datas
	return c.RenderJSON(&result)
}
