package manipulator

import (
	"github.com/go-xorm/xorm"
	kEmail "github.com/zuiwuchang/king-go/net/email"
	"store-web/app/db/data"
	"store-web/app/db/dberr"
	"store-web/app/log"
	"store-web/app/utils"
	"strings"
	"time"
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
func (u User) Register(host, email, pwd, code string) (nuser *data.User, e error) {
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
	user.Created = time.Now()
	user.LastEmail = user.Created

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
		// 發送 激活 郵件
		text, en := data.GetActiveCode(user.ID, user.Created.Unix())
		if en != nil {
			if log.Warn != nil {
				log.Warn.Println(en)
			}
			return
		}
		text, en = data.GetActiveEmail(
			&data.ActiveContext{
				Host:  host,
				Email: user.Email,
				ID:    user.ID,
				Code:  text,
			},
			systemInfo.ActiveText,
		)
		if en != nil {
			if log.Warn != nil {
				log.Warn.Println(en)
			}
			return
		}
		if en = kEmail.SendSSLEmail(
			systemInfo.SMTP,
			systemInfo.Email,
			systemInfo.Password,
			user.Email,
			systemInfo.ActiveTitle,
			text,
			kEmail.TypeHTML,
		); en != nil && log.Warn != nil {
			log.Warn.Println(en)
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
	user.UserGroup = utils.Separator + utils.GroupRootFlag + utils.Separator
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
	if systemInfo.Register == data.RegisterInvite {
		// 驗證 邀請碼
	}
	// 增加 用戶
	_, e = session.InsertOne(user)
	if e != nil {
		if log.Error != nil {
			log.Error.Println(e)
		}
		return
	}
	return
}

// Active 激活 帳號
func (User) Active(id int64, code string) (user *data.User, e error) {
	code = strings.TrimSpace(code)
	if code == "" {
		e = dberr.ErrUserCodeNotMatch
		if log.Warn != nil {
			log.Warn.Println(e, id, code)
		}
		return
	}
	if id == 0 {
		e = dberr.ErrUserNotFound
		if log.Warn != nil {
			log.Warn.Println(e, id)
		}
		return
	}

	// session
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

	// find
	bean := &data.User{ID: id}
	var ok bool
	if ok, e = session.Get(bean); e != nil {
		if log.Error != nil {
			log.Error.Println(e)
		}
		return
	} else if !ok || bean.Active {
		e = dberr.ErrUserNotFound
		if log.Warn != nil {
			log.Warn.Println(e, id)
		}
		return
	}

	// code
	var str string
	str, e = data.GetActiveCode(id, bean.Created.Unix())
	if e != nil {
		if log.Error != nil {
			log.Error.Println(e)
		}
		return
	}
	if code != str {
		e = dberr.ErrUserCodeNotMatch
		if log.Warn != nil {
			log.Warn.Println(e, id, code)
		}
		return
	}

	// active
	bean.Active = true
	_, e = session.ID(id).Cols(data.ColUserActive).Update(bean)
	if e != nil {
		if log.Error != nil {
			log.Error.Println(e)
		}
		return
	}

	user = bean
	return
}

// SendActiveEmail 發送 激活 郵件
func (User) SendActiveEmail(host string, user *data.User) (e error) {
	var mSys SystemInfo
	var systemInfo *data.SystemInfo
	systemInfo, e = mSys.Get()
	if e != nil {
		return
	}

	// 發送 激活 郵件
	text, en := data.GetActiveCode(user.ID, user.Created.Unix())
	if en != nil {
		if log.Warn != nil {
			log.Warn.Println(en)
		}
		return
	}
	text, en = data.GetActiveEmail(
		&data.ActiveContext{
			Host:  host,
			Email: user.Email,
			ID:    user.ID,
			Code:  text,
		},
		systemInfo.ActiveText,
	)
	if en != nil {
		if log.Warn != nil {
			log.Warn.Println(en)
		}
		return
	}
	if en = kEmail.SendSSLEmail(
		systemInfo.SMTP,
		systemInfo.Email,
		systemInfo.Password,
		user.Email,
		systemInfo.ActiveTitle,
		text,
		kEmail.TypeHTML,
	); en != nil && log.Warn != nil {
		log.Warn.Println(en)
	}

	// 更新 郵件 請求時間
	user.LastEmail = time.Now()
	Engine().Id(user.ID).Cols(data.ColUserLastEmail).Update(user)
	return
}
