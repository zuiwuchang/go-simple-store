package configure

// DB 數據庫 配置
type DB struct {
	// 數據庫 驅動
	Driver string
	// 數據庫 連接字符串
	Str string
	// 是否打印 執行的 sql 指令
	Show bool

	// 緩存 設置
	Cache DBCache
}

// DBCache 數據庫 緩存 設置
type DBCache struct {
	// 全局 緩存大小
	Size int
}
