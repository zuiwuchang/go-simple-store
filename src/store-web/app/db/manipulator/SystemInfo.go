package manipulator

import (
	"store-web/app/db/data"
	"store-web/app/db/dberr"
	"store-web/app/log"
)

// SystemInfo 系統表
type SystemInfo struct {
}

// Get 返回 系統設置
func (SystemInfo) Get() (bean *data.SystemInfo, e error) {
	find := &data.SystemInfo{
		ID: 1,
	}
	var ok bool
	if ok, e = Engine().Get(find); e != nil {
		if log.Error != nil {
			log.Error.Println(e)
		}
		return
	} else if ok {
		bean = find
	} else {
		e = dberr.ErrSystemInfoEmpty
		if log.Error != nil {
			log.Error.Println(e)
		}
	}
	return
}
