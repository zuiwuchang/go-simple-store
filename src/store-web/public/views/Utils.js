// 定義的 一些 常量
define(function () {
    'use strict';
    return {
        // pages 總頁數
        // page 當前 頁數
        // tag tag 數量
        // firstPage lastPage 顯示文本
        CreateNavPage: function (pages, page, tag, firstPage, lastPage) {
            let rs = [];
            if (pages < 2) { // 少於2頁 不想要 分頁標籤
                return rs;
            }

            // tag
            tag = Math.floor(tag / 2);
            if (tag < 2) {
                tag = 2;
            }

            // 計算 首頁
            let start = page - tag;
            if (start < 1) {
                start = 1;
            }

            // 需要 首頁
            if (start != 1) {
                rs.push({
                    Name: firstPage,
                    Page: 1,
                });
            }
            // 當前頁之前
            let val;
            for (let i = 0; i < tag + 1; i++) {
                val = start + i;
                let node = {
                    Name: val,
                    Page: val,
                };
                rs.push(node);
                if (val == page) {
                    node.Class = "active";
                    break;
                }
            }
            if (val >= pages) {
                return rs;
            }
            // 當前頁之後
            let val2;
            let last = true;
            for (let i = 0; i < tag; i++) {
                val2 = val + i + 1;
                let node = {
                    Name: val2,
                    Page: val2,
                };
                rs.push(node);
                if (val2 == page) {
                    node.Class = "active";
                }

                if (val2 == pages) {
                    last = false;
                    break;
                }
            }
            // 需要 尾頁
            if (last) {
                rs.push({
                    Name: lastPage,
                    Page: pages,
                });
            }
            return rs;
        },
    }
})