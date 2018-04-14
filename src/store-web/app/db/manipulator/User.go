package manipulator

import (
	"errors"
	"github.com/go-xorm/xorm"
	"store-web/app/db/data"
	"store-web/app/db/dberr"
	"store-web/app/log"
)

// User .
type User struct {
}

// IsEmailExists 返回 email 是否已經 存在
func (User) IsEmailExists(email string) (yes bool, e error) {
	if email == "" {
		return
	}
	find := &data.User{
		Email: email,
	}
	yes, e = Engine().Get(find)
	if e != nil {
		if log.Error != nil {
			log.Error.Println(e)
		}
	}
	return
}

// GetByEmail 使用 email 查詢
func (User) GetByEmail(email string) (fuser *data.User, e error) {
	// 驗證 用戶名
	user := &data.User{}
	e = user.SetEmail(email)
	if e != nil {
		if log.Warn != nil {
			log.Warn.Println(e)
		}
		return
	}

	// 查找用戶
	var ok bool
	if ok, e = Engine().Get(user); e != nil {
		if log.Error != nil {
			log.Error.Println(e)
		}
		return
	} else if ok {
		fuser = user
	}
	return
}

// Register 註冊 新用戶
func (u User) Register(email, pwd, code string) (nuser *data.User, e error) {
	// 驗證 用戶名 密碼
	user := &data.User{}
	e = user.SetEmail(email)
	if e != nil {
		if log.Warn != nil {
			log.Warn.Println(e)
		}
		return
	}
	user.SetPwd(pwd)
	if e != nil {
		if log.Warn != nil {
			log.Warn.Println(e)
		}
		return
	}

	// 創建 session
	var session *xorm.Session
	session, e = NewTransaction()
	if e != nil {
		if log.Error != nil {
			log.Error.Println(e)
		}
		return
	}
	defer func() {
		if e == nil {
			session.Commit()
		} else {
			session.Rollback()
		}
		session.Close()
	}()

	// 驗證 管理員是否 註冊
	systemInfo := &data.SystemInfo{ID: 1}
	var ok bool
	ok, e = session.Get(systemInfo)
	if e != nil {
		if log.Error != nil {
			log.Error.Print(e)
		}
		return
	} else if !ok {
		e = dberr.ErrSystemInfoEmpty
		if log.Error != nil {
			log.Error.Println(e)
		}
		return
	} else if systemInfo.InitRoot {
		if e = u.register(session, systemInfo, user, code); e == nil {
			nuser = user
		}
		return
	}
	if e = u.registerRoot(session, systemInfo, user); e == nil {
		nuser = user
	}
	return
}

func (User) registerRoot(session *xorm.Session, systemInfo *data.SystemInfo, user *data.User) (e error) {
	// 設置 root 組 激活狀態
	user.UserGroup = "1"
	user.Active = true
	// 增加 用戶
	_, e = session.InsertOne(user)
	if e != nil {
		if log.Error != nil {
			log.Error.Println(e)
		}
		return
	}

	// 禁用 root 自動註冊
	systemInfo.InitRoot = true
	_, e = session.ID(1).Cols(data.SystemInfoColInitRoot).Update(systemInfo)
	if e != nil {
		if log.Error != nil {
			log.Error.Println(e)
		}
	}
	return
}
func (User) register(session *xorm.Session, systemInfo *data.SystemInfo, user *data.User, code string) (e error) {
	e = errors.New("no code")
	return
}
