package data

import (
	"time"
)

// App 軟件包 定義
type App struct {
	// 軟件 標識
	ID int64 `xorm:"pk autoincr 'id'"`

	// 軟件名稱
	Name string `xorm:"unique"`

	// 描述信息
	Info string

	// 軟件官網
	URL string `xorm:"'url'"`

	// 原始作者 email (多個以 : 分隔)
	Author string

	// 初創 時間
	Created time.Time `xorm:"created"`
}
