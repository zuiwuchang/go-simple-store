package data

import (
	"time"
)

const (
	// ColInviteCodeID .
	ColInviteCodeID = "id"
	// ColInviteCodeCode .
	ColInviteCodeCode = "code"
)

// InviteCode 註冊 邀請碼
type InviteCode struct {
	ID      int64     `xorm:"pk autoincr 'id'"`
	Code    string    `xorm:"unique"`
	Created time.Time `xorm:"created"`
}
