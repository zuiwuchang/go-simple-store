package data

import (
	"errors"
	kStrings "github.com/zuiwuchang/king-go/strings"
	"regexp"
	"time"
)

const (
	// ColUserActive .
	ColUserActive = "active"
)

// ErrBadUserPwd .
var ErrBadUserPwd = errors.New("bad password")

// User 用戶
type User struct {
	// 用戶 標識
	ID int64 `xorm:"pk autoincr 'id'"`
	// 登入 用戶名
	Email string `xorm:"unique"`
	// 登入 密碼 sha512
	Pwd string

	// 帳號 是否 激活
	Active bool

	// 註冊 時間
	Created time.Time

	// 最後發送 激活 email 時間
	LastEmail time.Time

	// 所屬 用戶組id (多個以 : 分隔)
	UserGroup string
}

// SetEmail .
func (u *User) SetEmail(val string) (e error) {
	e = kStrings.MatchGMail(val)
	if e != nil {
		return
	}
	u.Email = val
	return
}

var regexpUserPwd, _ = regexp.Compile(`^[a-f0-9]+$`)

// SetPwd .
func (u *User) SetPwd(val string) (e error) {
	if len(val) != 128 {
		e = ErrBadUserPwd
		return
	}
	if !regexpUserPwd.MatchString(val) {
		e = ErrBadUserPwd
		return
	}

	u.Pwd = val
	return
}
