package data

const (
	// SystemInfoColInitRoot .
	SystemInfoColInitRoot = "init_root"
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
}
