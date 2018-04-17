package controllers

import (
	"github.com/revel/revel"
	"store-web/app/ajax"
	"store-web/app/db/manipulator"
	"store-web/app/utils"
	"time"
)

// Active .
type Active struct {
	*revel.Controller
}

// Index .
func (c Active) Index() revel.Result {
	user, e := readSession(c.Session)
	if e != nil {
		return c.RenderError(e)
	}
	email := user.Email
	interval := utils.ActiveEmailInterval
	wait := interval - int(time.Now().Sub(user.LastEmail)/time.Second)
	if wait < 0 {
		wait = 0
	} else if wait > interval {
		wait = interval
	}
	return c.Render(email, wait, interval)
}

// AjaxResend 重發激活郵件
func (c Active) AjaxResend() revel.Result {
	var result ajax.Result

	user, e := readSession(c.Session)
	if e != nil {
		result.Code = ajax.Error
		result.Emsg = e.Error()
		return c.RenderJSON(&result)
	}
	// 驗證 時間
	if time.Now().Sub(user.LastEmail) < utils.ActiveEmailInterval*time.Second {
		result.Code = ajax.Error
		result.Emsg = c.Message("actIndex.request active email busy")
		return c.RenderJSON(&result)
	}

	// 發送 email
	var mUser manipulator.User
	e = mUser.SendActiveEmail(c.Request.Host, user)
	if e != nil {
		result.Code = ajax.Error
		result.Emsg = e.Error()
		return c.RenderJSON(&result)
	}

	// 更新 session
	writeSession(c.Session, user)

	return c.RenderJSON(&result)
}
