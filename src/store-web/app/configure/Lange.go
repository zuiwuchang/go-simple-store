package configure

import (
	"strings"
)

// Lange 語言 設置
type Lange struct {
	// 默認 語言
	Default string

	// 允許訪問的 語言
	Locale []string

	keys map[string]bool
}

// Get 傳入語言環境 返回 支持的語言
func (l *Lange) Get(locale string) string {
	if ok, _ := l.keys[locale]; ok {
		return locale
	}
	return l.Default
}
func (l *Lange) format() {
	l.Default = strings.TrimSpace(l.Default)
	if l.Default == "" {
		l.Default = "zh-TW"
	} else {
		str := strings.ToLower(l.Default)
		if str == "zh-cn" || str == "ru" || strings.HasPrefix(str, "ru-") {
			l.Default = "zh-TW"
		}
	}

	keys := make(map[string]bool)
	keys[l.Default] = true
	for i := 0; i < len(l.Locale); i++ {
		k := strings.TrimSpace(l.Locale[i])
		if k == "" {
			continue
		} else {
			str := strings.ToLower(k)
			if str == "zh-cn" || str == "ru" || strings.HasPrefix(str, "ru-") {
				continue
			}
		}

		keys[k] = true
	}
	l.keys = keys
}
