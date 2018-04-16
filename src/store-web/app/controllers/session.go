package controllers

import (
	"fmt"
	"store-web/app/db/data"
	"store-web/app/log"
	"store-web/app/utils"
	"strconv"
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
