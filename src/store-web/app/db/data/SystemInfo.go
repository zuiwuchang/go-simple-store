package data

import (
	"html/template"
	"store-web/app/db/dberr"
	"store-web/app/utils"
	"strings"
)

const (
	// SystemInfoColInitRoot .
	SystemInfoColInitRoot = "init_root"
	// SystemInfoColRegister .
	SystemInfoColRegister = "register"
	// SystemInfoColSMTP .
	SystemInfoColSMTP = "smtp"
	// SystemInfoColEmail .
	SystemInfoColEmail = "email"
	// SystemInfoColPassword .
	SystemInfoColPassword = "password"
	// SystemInfoColActiveTitle .
	SystemInfoColActiveTitle = "active_title"
	// SystemInfoColActiveText .
	SystemInfoColActiveText = "active_text"
)
const (
	// RegisterDisabled 關閉註冊
	RegisterDisabled = iota
	// RegisterOpen 開放註冊
	RegisterOpen
	// RegisterInvite 邀請註冊
	RegisterInvite
)

// SystemInfo 系統設置
type SystemInfo struct {
	ID int64 `xorm:"pk autoincr 'id'"`

	// 管理員 是否已註冊
	InitRoot bool

	// 註冊模式
	Register int

	// smtp 地址
	SMTP string `xorm:"'smtp'"`

	// 發送 email 用戶
	Email string
	// 發送 email 密碼
	Password string

	// 激活郵件 標題
	ActiveTitle string
	// 激活郵件 內容
	ActiveText string `xorm:"TEXT"`
}

// SetRegister .
func (s *SystemInfo) SetRegister(val int) (e error) {
	if val < RegisterDisabled || val > RegisterInvite {
		e = dberr.ErrSystemInfoUnknowRegisterMode
		return
	}
	s.Register = val
	return
}

// SetActiveTitle .
func (s *SystemInfo) SetActiveTitle(val string) (e error) {
	val = strings.TrimSpace(val)
	if val == "" {
		e = dberr.ErrSystemInfoActiveTitleEmpty
		return
	}
	s.ActiveTitle = val
	return
}

// SetActiveText .
func (s *SystemInfo) SetActiveText(val string) (e error) {
	val = strings.TrimSpace(val)
	if val == "" {
		e = dberr.ErrSystemInfoActiveTextEmpty
		return
	}

	t := template.New("active email")
	t, e = t.Parse(val)
	if e != nil {
		return
	}

	e = t.Execute(utils.NullWriter{}, &ActiveContext{
		Host:  "test.king.xxx",
		Email: "testName@xxx.xxx",
		ID:    0,
		Code:  "19111010",
	})
	if e != nil {
		return
	}

	s.ActiveText = val
	return
}
