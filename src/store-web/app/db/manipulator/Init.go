package manipulator

import (
	"github.com/go-xorm/xorm"
	"store-web/app/configure"
	"store-web/app/db/data"
	"store-web/app/log"
)

// Logger .
var Logger = log.Logger
var _Engine *xorm.Engine

// Init 初始化 數據庫
func Init() {
	db := configure.Get().DB

	// 初始化 數據庫 引擎
	engine, e := xorm.NewEngine(
		db.Driver,
		db.Str,
	)
	if e != nil {
		Logger.Fault.Fatalln(e)
	}
	e = engine.Ping()
	if e != nil {
		Logger.Fault.Fatalln(e)
	}

	// show sql
	if db.Show {
		engine.ShowSQL(true)
	}

	// 保存 單件
	_Engine = engine

	// 初始化 表
	initDB()

	// 初始化 緩存
	initCache()

}

// Engine .
func Engine() *xorm.Engine {
	return _Engine
}

// NewSession .
func NewSession() *xorm.Session {
	return _Engine.NewSession()
}

// NewTransaction 創建 一個 事務
func NewTransaction() (session *xorm.Session, e error) {
	session = _Engine.NewSession()
	e = session.Begin()
	if e != nil {
		session.Close()
		return
	}
	return
}
func initCache() {
	cache := configure.Get().DB.Cache
	engine := Engine()
	var e error

	// system
	{
		cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1)
		e = engine.MapCacher(&data.SystemInfo{}, cacher)
		if e != nil {
			Logger.Fault.Fatalln(e)
		}
	}

	// default
	if cache.Size > 0 {
		cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), cache.Size)
		engine.SetDefaultCacher(cacher)
	}
}
func initDB() {
	session, e := NewTransaction()
	if e != nil {
		Logger.Fault.Fatalln(e)
	}
	defer func() {
		if e == nil {
			session.Commit()
			session.Close()
		} else {
			session.Rollback()
			session.Close()

			Logger.Fault.Fatalln(e)
		}
	}()

	if e = initTable(
		session,
		&data.SystemInfo{},
		&data.App{},
		&data.AppVersion{},
		&data.User{},
		&data.UserGroup{},
	); e != nil {
		return
	}
	// 初始化 系統表
	if e = initSystemInfo(session); e != nil {
		return
	}

	// 初始化 root 組
	if e = initGroup(session); e != nil {
		return
	}
}
func initSystemInfo(session *xorm.Session) (e error) {
	bean := &data.SystemInfo{
		ID: 1,
	}
	var ok bool
	if ok, e = session.Get(bean); e != nil {
		return
	} else if !ok {
		_, e = session.Insert(bean)
	}
	return
}
func initGroup(session *xorm.Session) (e error) {
	bean := &data.UserGroup{
		ID: 1,
	}
	var ok bool
	if ok, e = session.Get(bean); e != nil {
		return
	} else if !ok {
		// 創建 root 組
		bean.Name = "root"
		_, e = session.Insert(bean)
	}
	return
}
func initTable(session *xorm.Session, beans ...interface{}) (e error) {
	var ok bool
	for _, bean := range beans {
		if ok, e = session.IsTableExist(bean); e != nil {
			return
		} else if ok {
			// 同步 表
			session.Sync2(bean)
		} else {
			// 創建 表
			if e = session.CreateTable(bean); e != nil {
				return
			} else if e = session.CreateIndexes(bean); e != nil {
				return
			} else if e = session.CreateUniques(bean); e != nil {
				return
			}
		}
	}
	return
}
