package data

import (
	"time"
)

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
	Created time.Time `xorm:"created"`

	// 最後發送 激活 email 時間
	LastEmail time.Time

	// 所屬 用戶組id (多個以 : 分隔)
	UserGroup string
}
