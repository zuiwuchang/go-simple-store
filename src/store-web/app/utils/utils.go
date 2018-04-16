package utils

import (
	"strings"
)

// HasFlags 返回 是否存在 指定 標記
func HasFlags(flags, flag string) (yes bool) {
	flags = strings.TrimSpace(flags)
	if flags == "" {
		return
	}
	flag = strings.TrimSpace(flag)
	if flag == "" {
		return
	}
	if !strings.HasPrefix(flags, Separator) {
		flags = Separator + flags
	}
	if !strings.HasSuffix(flags, Separator) {
		flags += Separator
	}

	yes = strings.Index(flags, flag) != -1
	return
}
