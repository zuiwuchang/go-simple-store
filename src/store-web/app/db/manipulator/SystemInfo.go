package manipulator

import (
	"store-web/app/db/data"
	"store-web/app/db/dberr"
	"store-web/app/log"
)

// SystemInfo 系統表
type SystemInfo struct {
}

// Get 返回 系統設置
func (SystemInfo) Get() (bean *data.SystemInfo, e error) {
	find := &data.SystemInfo{
		ID: 1,
	}
	var ok bool
	if ok, e = Engine().Get(find); e != nil {
		if log.Error != nil {
			log.Error.Println(e)
		}
		return
	} else if ok {
		bean = find
	} else {
		e = dberr.ErrSystemInfoEmpty
		if log.Error != nil {
			log.Error.Println(e)
		}
	}
	return
}

// SaveRegister 修改 註冊模式
func (SystemInfo) SaveRegister(registerMode int) (e error) {
	var systemInfo data.SystemInfo
	e = systemInfo.SetRegister(registerMode)
	if e != nil {
		if log.Warn != nil {
			log.Warn.Println(e, registerMode)
		}
		return
	}
	_, e = Engine().ID(1).Cols(data.SystemInfoColRegister).Update(&systemInfo)
	if e != nil {
		if log.Error != nil {
			log.Error.Println(e)
		}
	}
	return
}

// SaveSMTP 修改 smtp 配置
func (SystemInfo) SaveSMTP(smtp, email, pwd string) (e error) {
	systemInfo := data.SystemInfo{
		SMTP:     smtp,
		Email:    email,
		Password: pwd,
	}
	_, e = Engine().
		ID(1).
		Cols(
			data.SystemInfoColSMTP,
			data.SystemInfoColEmail,
			data.SystemInfoColPassword,
		).
		Update(systemInfo)
	if e != nil {
		if log.Error != nil {
			log.Error.Println(e)
		}
	}
	return
}

// SaveActive 存儲 激活郵件 配置
func (SystemInfo) SaveActive(title, text string) (e error) {
	var systemInfo data.SystemInfo
	e = systemInfo.SetActiveTitle(title)
	if e != nil {
		if log.Warn != nil {
			log.Warn.Println(e, title)
		}
		return
	}
	e = systemInfo.SetActiveText(text)
	if e != nil {
		if log.Warn != nil {
			log.Warn.Println(e, text)
		}
		return
	}

	_, e = Engine().
		ID(1).
		Cols(data.SystemInfoColActiveTitle, data.SystemInfoColActiveText).
		Update(&systemInfo)
	if e != nil {
		if log.Error != nil {
			log.Error.Println(e)
		}
	}
	return
}
