export as namespace strings;

declare module 'king/strings' {
    /*~ 驗證 字符串 是否是 ipv4 格式 xxx.xxx.xxx.xxx */
    function IsIPV4(str: string): string;

    enum MatchGMailRS {
        //匹配 成功
        Success,
    
        //匹配 用戶名@主機名
        SplitLess, //沒有以 @ 分隔 用戶名 主機名
        SplitLarge,	//出現 多個 @
    
        //匹配用戶名
        LenLess, 	//用戶名太短
        LenLarge, 	//用戶名太長
        BadBegin,	//驗證用戶名 開始
        BadEnd,		//驗證用戶名 結尾
        PointLink,	//多個 . 相連
    
        //匹配主機名
        BadHost,
    };
    /*
        驗證 字符串是否符合 gmail 格式 要求

        用戶名@主機名
        用戶名
            字符[a-zA-Z0-9]長度爲 [6,30] 
            以 [a-zA-Z0-9] 開始 結束
            中間只能是 [a-zA-Z0-9\.] 且 . 和 . 不能相連
        域名
            以 .[a-zA-Z]{2,} 結尾
            不能以 . 開頭 且 . 不能相連
            由 [a-zA-Z0-9\-_\.]組成
    */
    function MatchGMail(str: string): MatchGMailRS;

    //將字符串中的 html 保留字 轉義
    function HtmlEncode(str:string):string;
    //將 HtmlEncode 解碼
    function HtmlDecode(str:string):string;

    /*
        將日期 Date 轉化爲指定格式的 字符串

        fmt yyyy-mm-dd HH:MM:SS
    */
    function FormatDate(d:Date,fmt:string):string;
}