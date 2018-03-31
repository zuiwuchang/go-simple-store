package configure

// Log 日誌 配置
type Log struct {
	// 顯示 短 源文件名
	Short bool
	// 顯示 長 源文件名
	Long bool
	// 需要顯示的 日誌
	Logs []string
}
