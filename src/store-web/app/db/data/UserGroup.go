package data

// UserGroup 用戶組 定義 每個 組定義了 可以 如何 操作 軟件包
type UserGroup struct {
	// 組 標識 如果 ID 為 1 則為 root 組
	ID int64 `xorm:"pk autoincr 'id'"`

	// 組名
	Name string `xorm:"unique"`

	// 可維護的 軟件包 id (多個以 : 分隔)
	AppMaintainer string
	// 可下載的 軟件包 id (多個以 : 分隔)
	AppUser string
}
