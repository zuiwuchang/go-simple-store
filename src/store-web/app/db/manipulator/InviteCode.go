package manipulator

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/google/uuid"
	"store-web/app/db/data"
	"store-web/app/log"
	"strings"
)

// InviteCode 邀請碼 管理
type InviteCode struct {
}

// Code 創建 邀請碼
func (InviteCode) Code() (code string, e error) {
	var u uuid.UUID
	//創建一個 uuid
	u, e = uuid.NewUUID()
	if e != nil {
		if log.Warn != nil {
			log.Warn.Println(e)
		}
		return
	}
	b := u[:]

	// md5
	enc := md5.New()
	var n int
	n, e = enc.Write(b)
	if e != nil {
		if log.Error != nil {
			log.Error.Println(e)
		}
		return
	} else if n != len(b) {
		e = errors.New("busy write")
		if log.Error != nil {
			log.Error.Println(e)
		}
		return
	}
	str := hex.EncodeToString(enc.Sum(nil))

	// 寫入 數據庫
	bean := &data.InviteCode{
		Code: str,
	}
	_, e = Engine().InsertOne(bean)
	if e != nil {
		if log.Error != nil {
			log.Error.Println(e)
		}
		return
	}

	// 返回
	code = str
	return
}

func (InviteCode) findSession(session *xorm.Session, code string) *xorm.Session {
	code = strings.TrimSpace(code)
	if code != "" {
		session = session.Where(
			fmt.Sprintf(`(%s like ?)`, data.ColInviteCodeCode),
			code,
		)
	}
	return session
}

// Count .
func (m InviteCode) Count(code string) (n int64, e error) {
	session := m.findSession(NewSession(), code)
	defer session.Close()

	n, e = session.Count(&data.InviteCode{})
	if e != nil {
		if log.Error != nil {
			log.Error.Println(e)
		}
	}
	return
}

// Find .
func (m InviteCode) Find(start, rows int64, code string) (datas []data.InviteCode, e error) {
	var beans []data.InviteCode

	session := m.findSession(NewSession(), code)
	defer session.Close()

	e = session.Limit(int(rows), int(start)).Desc(data.ColInviteCodeID).Find(&beans)
	if e != nil {
		if log.Error != nil {
			log.Error.Println(e)
		}
		return
	}
	datas = beans
	return
}

// Remove .
func (m InviteCode) Remove(id int64) (e error) {
	if id == 0 {
		return
	}

	_, e = Engine().Id(id).Delete(&data.InviteCode{})
	if e != nil {
		if log.Error != nil {
			log.Error.Println(e)
		}
	}
	return
}
