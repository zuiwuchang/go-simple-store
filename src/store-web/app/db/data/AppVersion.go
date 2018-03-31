package data

import (
	"time"
)

// AppVersion 軟件包的 一個 可使用 版本
type AppVersion struct {
	// 唯一標識 包 sha512
	ID string `xorm:"pk 'id'"`
	// 所屬的 軟件包
	App int64 `xorm:"index"`

	// 當前 版本 以utf8字符串 < 判斷 版本 大小
	Version string
	// 當前 版本 描述
	VersionInfo string

	// 上傳時間
	Upload time.Time `xorm:"created"`

	// 包創建 時間
	Created time.Time

	// 當前 維護者 (多個以 : 分隔)
	Maintainer string
	// 當前 開放人員 (多個以 : 分隔)
	Author string
}
