{
    // 語言配置 使用 i18n (注意只支持 i18n 中 人類的語言)
    Lange:{
        // 默認語言 默認值為 zh-TW
        Default:"zh-TW",
        // 已經翻譯好的 語言
        Locale:[
            "zh-TW",
        ],
    },
    // 數據庫配置
    DB:{
        // 數據庫 驅動
        Driver:"mysql",
        // 數據庫 連接字符串
        Str:"GoSimpleStore:12345678@/GoSimpleStore?charset=utf8",
        // 是否打印 執行的 sql 指令
        Show:true,

        // 緩存 設置
        Cache:{
            // 全局 緩存大小
            Size:1000,
        },
    },
    // 日誌 配置
    Log:{
         // 顯示 短 源文件名
        Short:true,
        // 顯示 長 源文件名
        Long:false,
        // 需要顯示的 日誌
        Logs:[
            "trace",
            "debug",
            "info",
            "warn",
            "error",
            "fault",
        ],        
    },
   
}