package controllers

import (
	"errors"
	"fmt"
	"store-web/app/db/data"
	"store-web/app/log"
	"store-web/app/utils"
	"strconv"
	"time"
)

const (
	// SessionKeyID .
	SessionKeyID = utils.SessionKeyID
	// SessionKeyEmail .
	SessionKeyEmail = utils.SessionKeyEmail
	// SessionKeyActive .
	SessionKeyActive = utils.SessionKeyActive
	// SessionKeyCreated .
	SessionKeyCreated = utils.SessionKeyCreated
	// SessionKeyLastEmail .
	SessionKeyLastEmail = utils.SessionKeyLastEmail
	// SessionKeyGroup .
	SessionKeyGroup = utils.SessionKeyGroup
)

// ErrSessionEmpty .
var ErrSessionEmpty = errors.New("session empty")

func writeSession(session map[string]string, user *data.User) {
	session[SessionKeyID] = fmt.Sprint(user.ID)
	session[SessionKeyEmail] = user.Email
	if user.Active {
		session[SessionKeyActive] = "1"
	} else {
		session[SessionKeyActive] = "0"
	}
	session[SessionKeyCreated] = fmt.Sprint(user.Created.Unix())
	session[SessionKeyLastEmail] = fmt.Sprint(user.LastEmail.Unix())
	session[SessionKeyGroup] = user.UserGroup
}
func readSession(session map[string]string) (user *data.User, e error) {
	var rs data.User
	// id
	str, ok := session[SessionKeyID]
	if !ok {
		e = ErrSessionEmpty
		return
	}
	rs.ID, e = strconv.ParseInt(str, 10, 64)
	if e != nil {
		return
	}
	// email
	rs.Email, _ = session[SessionKeyEmail]
	//active
	str, _ = session[SessionKeyActive]
	if str == "1" {
		rs.Active = true
	}

	// created
	var unix int64
	str, _ = session[SessionKeyCreated]
	unix, e = strconv.ParseInt(str, 10, 64)
	if e != nil {
		return
	}
	rs.Created = time.Unix(unix, 0)

	// last email
	str, _ = session[SessionKeyLastEmail]
	unix, e = strconv.ParseInt(str, 10, 64)
	if e != nil {
		return
	}
	rs.LastEmail = time.Unix(unix, 0)

	// group
	rs.UserGroup, _ = session[SessionKeyGroup]

	user = &rs
	return
}
func isLogin(session map[string]string) bool {
	if key, ok := session[SessionKeyID]; ok {
		if id, e := strconv.ParseInt(key, 10, 64); e != nil {
			if log.Error != nil {
				log.Error.Println(e, key)
			}
			return false
		} else if id == 0 {
			if log.Error != nil {
				log.Error.Println("id == 0")
			}
			return false
		}

		return true
	}
	return false
}
