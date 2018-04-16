package data

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/zuiwuchang/king-go/strings"
	kStrings "github.com/zuiwuchang/king-go/strings"
	"html/template"
)

// ActiveContext 激活 郵件 上下文
type ActiveContext struct {
	// 用戶 訪問 地址
	Host string

	// 註冊 用戶名
	Email string

	// 註冊 ID
	ID int64

	// 激活碼
	Code string
}

// GetActiveEmail .
func GetActiveEmail(context *ActiveContext, text string) (str string, e error) {
	t := template.New("active email")
	t, e = t.Parse(text)
	if e != nil {
		return
	}

	var w bytes.Buffer
	e = t.Execute(&w, context)
	if e != nil {
		return
	}
	str = strings.BytesToString(w.Bytes())
	return
}

// GetActiveCode 返回 激活碼
func GetActiveCode(id int64, unixCreated int64) (code string, e error) {
	str := fmt.Sprint(id, unixCreated)
	b := kStrings.StringToBytes(str)

	enc := md5.New()
	var n int
	n, e = enc.Write(b)
	if e != nil {
		return
	} else if n != len(b) {
		e = errors.New("busy write")
		return
	}

	code = hex.EncodeToString(enc.Sum(nil))
	return
}
