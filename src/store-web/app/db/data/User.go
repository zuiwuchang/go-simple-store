package data

// User 用戶
type User struct {
	// 用戶 標識
	ID int64 `xorm:"pk autoincr 'id'"`
	// 登入 用戶名
	User string `xorm:"unique"`
	// 登入 密碼 sha512
	Pwd string

	// 所屬 用戶組id (多個以 : 分隔)
	UserGroup string
}
