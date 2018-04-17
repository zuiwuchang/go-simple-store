define(["Utils", "angular"], function () {
    'use strict';
    let Utils = require("Utils");
    return function (context) {
        let App = context.App;
        let Lange = context.Lange;

        // 創建 angular 模塊
        let app = angular.module('app', []);
        app.config(function ($httpProvider) {
            $httpProvider.defaults.transformRequest = function (data) {
                if (data === undefined) {
                    return data;
                }
                return $.param(data);
            }
            $httpProvider.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded; charset=UTF-8';
        });

        app.controller("ctrl-new",
            [
                "$scope", "$http",
                function ($scope, $http) {
                    $scope.disabled = false;
                    $scope.error = "";
                    $scope.code = "";

                    $scope.onCreate = function () {
                        // 禁用 ui
                        $scope.disabled = true;
                        $scope.error = "";
                        $scope.code = "";

                        // 請求測試
                        $http({
                            method: 'get',
                            url: '/Root/AjaxCode',
                        }).then((response) => {
                            let rs = response.data;
                            if (rs.Code) {
                                $scope.error = rs.Emsg;
                            } else {
                                $scope.code = rs.Str;
                            }
                        }, (response) => {
                            $scope.error = response.data.description;
                        }).finally(() => {
                            // 啟用 ui
                            $scope.disabled = false;
                        });
                    };
                },
            ]
        );
        app.controller("ctrl-search",
            [
                "$scope", "$http", "$log",
                function ($scope, $http, $log) {
                    const ActionCode = 0;
                    const ActionRows = 1;

                    $scope.disabled = false;
                    $scope.error = "";
                    $scope.action = ActionCode;
                    $scope.show = false;

                    $scope.code = "";
                    $scope.rows = 20;
                    $scope.pages = [];

                    $scope.title = Lange["f.code"];
                    $scope.onBtnCode = function () {
                        if ($scope.action != ActionCode) {
                            $scope.action = ActionCode;
                            $scope.title = Lange["f.code"];
                        }
                    };
                    $scope.onBtnRows = function () {
                        if ($scope.action != ActionRows) {
                            $scope.action = ActionRows;
                            $scope.title = Lange["f.rows"];
                        }
                    };
                    $scope.onBtnReset = function () {
                        $scope.code = "";
                        $scope.rows = 20;
                        $scope.error = "";
                    };
                    let cache = {
                        Code: null,
                        Rows: null,
                        Page: null,
                        Pages: null,
                    };

                    function updateView(rs, rows, pages, page, code) {
                        // 更新 緩存
                        cache.Code = code;
                        cache.Rows = rows;
                        cache.Page = page;
                        cache.Pages = pages;

                        // 更新 翻頁按鈕
                        $scope.pages = Utils.CreateNavPage(rs.Pages, page, 5, Lange["p.f"], Lange["p.l"])
                    }
                    function doSearch(rows, pages, page, code) {
                        // 禁用 ui
                        $scope.disabled = true;
                        $scope.error = "";
                        $scope.show = false;

                        // 請求 數據
                        $http({
                            method: 'post',
                            url: '/Root/AjaxFindCode',
                            data: {
                                code: code,
                                rows: rows,
                                page: page,
                                pages: pages,
                            },
                        }).then((response) => {
                            let rs = response.data;
                            if (rs.Code) {
                                $scope.error = rs.Emsg;
                            } else {
                                updateView(rs, rows, pages, page, code);
                                $scope.show = true;
                            }
                        }, (response) => {
                            $scope.error = response.data.description;
                        }).finally(() => {
                            // 啟用 ui
                            $scope.disabled = false;
                        });
                    }
                    $scope.onBtnSearch = function () {
                        // 檢查是否需要更新
                        if ($scope.rows == cache.Rows && $scope.code == cache.Code && 1 == cache.Page) {
                            $log.info("not need search");
                            return;
                        }
                        doSearch($scope.rows, 0, 1, $scope.code);
                    };
                    $scope.onBtnPage = function (page) {
                        // 檢查是否需要更新
                        if (page == cache.Page) {
                            $log.info("not need search");
                            return;
                        }
                        doSearch(cache.Rows, cache.Pages, page, cache.Code);
                    }
                },
            ]
        );

        // 運行 angular 模塊
        angular.bootstrap(App, ['app']);
    };
});