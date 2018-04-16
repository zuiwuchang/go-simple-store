package data

import (
	"bytes"
	"github.com/zuiwuchang/king-go/strings"
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
