package configure

import (
	"encoding/json"
	"github.com/google/go-jsonnet"
	kStrings "github.com/zuiwuchang/king-go/strings"
	"io/ioutil"
)

// Configure 服務器 配置
type Configure struct {
	// 版本 信息 由 git 自動創建
	Version string

	// 數據庫 配置
	DB DB `json:"DB"`

	// 日誌 配置
	Log Log
	// 語言設置
	Lange Lange
}

var gConfigure Configure

// Get 返回 全局的 配置 單件
func Get() *Configure {
	return &gConfigure
}

// Init 初始化
func Init(filename string) {
	// read file
	b, e := ioutil.ReadFile(filename)
	if e != nil {
		panic(e)
	}

	// jsonnet to json
	vm := jsonnet.MakeVM()
	str, e := vm.EvaluateSnippet("", kStrings.BytesToString(b))
	if e != nil {
		panic(e)
	}
	b = kStrings.StringToBytes(str)

	// json to go struct
	e = json.Unmarshal(b, &gConfigure)
	if e != nil {
		panic(e)
	}

	// 標準化 配置
	gConfigure.Lange.format()

}
