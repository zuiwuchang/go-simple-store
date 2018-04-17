package ajax

import (
	"store-web/app/db/data"
)

// Result ajax 請求返回結果
type Result struct {
	// 錯誤碼 為0代表 成功
	Code int
	// 錯誤描述
	Emsg string

	// int64 返回值
	Value int64
	// string 返回值
	Str string
}

// ResultFindCode .
type ResultFindCode struct {
	Result
	Pages int64
	Data  []*data.InviteCode
}
