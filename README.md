# go-simple-store
使用 go 實現的一個 簡陋的 軟件倉庫

# what
這是 一個 簡陋的 軟件倉庫 以便與 能夠將 發佈給客戶的 軟體 進行自動 升級 管理 其中並沒有 解決軟體依賴等高級 特性(只是 為軟體提供了 一個版本 信息 以及 自動 更新 策略)

寫這個 工具 的目的 只是 受不了 所在公司 頻繁的 產品修改後 又要 孤手動去 為 用戶 升級(孤又不是 廉價勞動力 fucking) 

這個項目的 本意 是解決 所在公司 windows下的 產品的發佈 但應該也能在 linux下 工作 (好吧 其實開發測試 都是在mint下 做的 奈何寄人籬下 不得不被windows折磨)

如果你的客戶在 linux下 太恭喜你了 你最好使用 yum 或 deb 等更成熟的解決方案 但 也不妨一試本項目 畢竟 這個項目 比 yum deb 等簡單許多 且可以同時工作在 red hat 和 debian 系列的 系統上

# web
實現了 軟件倉庫 的 服務器 功能 以及 開放了一個 web 的管理 接口

# utils
提供了 將 軟件 打包 和一些 其它相關 工具

# system
安裝到 客戶處的 軟件倉庫 系統

