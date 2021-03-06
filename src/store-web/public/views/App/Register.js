define(["angular", "Const", "king/strings", "crypto-js/sha512"], function () {
    'use strict';
    let strings = require("king/strings");
    let Const = require("Const");
    let Register = Const.Register;
    let sha512 = require("crypto-js/sha512");
   
    let App;
    let Lange;
    function runApp(isInvite) {
        const CSSHasNone = "";
        const CSSHasError = "has-error";
        const CSSHasSuccess = "has-success";

        // 創建 angular 模塊
        let app = angular.module('app', []);
        app.config(["$httpProvider", function ($httpProvider) {
            $httpProvider.defaults.transformRequest = function (data) {
                if (data === undefined) {
                    return data;
                }
                return $.param(data);
            }
            $httpProvider.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded; charset=UTF-8';
        }]);

        // 數據共享 郵箱/密碼/邀請碼
        let Email = null;
        let Password = null;
        let Code = null;

        app.controller("form",
            [
                "$scope", "$rootScope", "$log", "$http",
                function ($scope, $rootScope, $log, $http) {
                    $scope.disabled = false;
                    $scope.error = ""
                    $scope.onSubmit = function () {
                        if (!Email) {
                            $log.info("Email has some error")
                            return;
                        } else if (!Password) {
                            $log.info("Password has some error")
                            return;
                        } else if (!Code && isInvite) {
                            $log.info("Code has some error")
                            return;
                        }

                        // 禁用 ui
                        $rootScope.$broadcast('disabledChange', true);
                        $scope.disabled = true;
                        let doError = (emsg) => {
                            $scope.error = emsg;
                        };
                        $http({
                            method: 'post',
                            url: '/App/AjaxRegister',
                            data: {
                                email: Email,
                                pwd: sha512(Password).toString(),
                                code: Code,
                            },
                        }).then((response) => {
                            let rs = response.data;
                            if (rs.Code) {
                                doError(rs.Emsg);
                            } else if (rs.Value) {
                                // 成功 轉到首頁
                                window.location.href = "/";
                            } else {
                                doError(rs.Emsg)
                            }
                        }, (response) => {
                            doError(response.data.description)
                        }).finally(() => {
                            // 啟用 ui
                            $rootScope.$broadcast('disabledChange', false);
                            $scope.disabled = false;
                        });
                    }
                }
            ]
        );
        app.controller("ctrl-email",
            [
                "$scope","$rootScope", "$http",
                function ($scope,$rootScope, $http) {
                    $scope.email = "";
                    $scope.error = "";
                    $scope.disabled = false;
                    $scope.$on("disabledChange", function (evt, yes) {
                        $scope.disabled = yes;
                    });
                    $scope.onChange = function () {
                        // 重置 email 驗證
                        Email = null;
                        $scope.inputClass = CSSHasNone;
                        $scope.error = "";
                    };
                    // 緩存 錯誤用戶名
                    let cacheBadEmail = {};
                    $scope.onBlur = function () {
                        // 獲取 用戶名
                        let email = $scope.email.trim();

                        // 查詢緩存
                        let cache = cacheBadEmail[email];
                        if (cache) {
                            $scope.inputClass = CSSHasError;
                            $scope.error = cache;
                            return;
                        }

                        // 驗證用戶名
                        let rs = strings.MatchGMail(email);
                        if (rs) {
                            $scope.inputClass = CSSHasError;
                            let matchRS = strings.MatchGMailRS;
                            switch (rs) {
                                case matchRS.SplitLess:
                                    cacheBadEmail[email] = Lange["E.GM.SplitLess"];
                                    $scope.error = Lange["E.GM.SplitLess"];
                                    break;
                                case matchRS.SplitLarge:
                                    cacheBadEmail[email] = Lange["E.GM.SplitLarge"];
                                    $scope.error = Lange["E.GM.SplitLarge"];
                                    break;

                                case matchRS.LenLess:
                                    cacheBadEmail[email] = Lange["E.GM.LenLess"];
                                    $scope.error = Lange["E.GM.LenLess"];
                                    break;
                                case matchRS.LenLarge:
                                    cacheBadEmail[email] = Lange["E.GM.LenLarge"];
                                    $scope.error = Lange["E.GM.LenLarge"];
                                    break;
                                case matchRS.BadBegin:
                                    cacheBadEmail[email] = Lange["E.GM.BadBegin"];
                                    $scope.error = Lange["E.GM.BadBegin"];
                                    break;
                                case matchRS.BadEnd:
                                    cacheBadEmail[email] = Lange["E.GM.BadEnd"];
                                    $scope.error = Lange["E.GM.BadEnd"];
                                    break;
                                case matchRS.PointLink:
                                    cacheBadEmail[email] = Lange["E.GM.PointLink"];
                                    $scope.error = Lange["E.GM.PointLink"];
                                    break;
                                case matchRS.BadHost:
                                    cacheBadEmail[email] = Lange["E.GM.BadHost"];
                                    $scope.error = Lange["E.GM.BadHost"];
                                    break;
                                default:
                                    cacheBadEmail[email] = Lange["E.Unknow"];
                                    $scope.error = Lange["E.Unknow"];
                            }
                            return;
                        }

                        // 服務器 驗證
                        $scope.inputClass = CSSHasNone;
                        let doError = (emsg) => {
                            $scope.inputClass = CSSHasError;
                            $scope.error = emsg;
                        };
                        $http({
                            method: 'post',
                            url: '/App/AjaxIsEmailExists',
                            data: {
                                email: email,
                            },
                        }).then((response) => {
                            let rs = response.data;
                            if (rs.Code) {
                                $scope.inputClass = CSSHasError;
                                $scope.error = rs.Emsg;
                            } else if (rs.Value) {
                                doError(Lange["e.email exists"])
                            } else {
                                $scope.inputClass = CSSHasSuccess;
                                // 設置 數據成功
                                Email = $scope.email;

                                $rootScope.$broadcast('emailOk');
                            }
                        }, (response) => {
                            doError(response.data.description)
                        });
                    };
                }
            ]
        );
        app.controller("ctrl-password",
            [
                "$scope",
                function ($scope) {
                    $scope.pwd = "";
                    $scope.error = "";
                    $scope.disabled = false;
                    $scope.$on("disabledChange", function (evt, yes) {
                        $scope.disabled = yes;
                    });
                    $scope.$on("emailOk", function (evt) {
                        $scope.onBlur();
                    });
                    $scope.onChange = function () {
                        // 重置 password 驗證
                        Password = null;
                        $scope.inputClass = CSSHasNone;
                        $scope.error = "";
                    };
                    $scope.onBlur = function () {
                        // 獲取 密碼
                        let val = $scope.pwd;

                        // 驗證 密碼
                        if (!val) {
                            $scope.inputClass = CSSHasError;
                            $scope.error = Lange["e.pwd empty"];
                            return
                        }

                        // 設置 成功
                        $scope.inputClass = CSSHasSuccess;
                        Password = val;
                    }
                }
            ]
        );
        // 需要 邀請碼
        if (isInvite) {
            app.controller("ctrl-invite",
                [
                    "$scope",
                    function ($scope) {
                        $scope.code = "";
                        $scope.error = "";
                        $scope.disabled = false;
                        $scope.$on("disabledChange", function (evt, yes) {
                            $scope.disabled = yes;
                        });
                        $scope.onChange = function () {
                            // 重置 Code 驗證
                            Code = null;
                            $scope.inputClass = CSSHasNone;
                            $scope.error = "";
                        };
                        $scope.onBlur = function () {
                            // 獲取 驗證
                            let val = $scope.code;

                            // 驗證 驗證
                            if (!val) {
                                $scope.inputClass = CSSHasError;
                                $scope.error = Lange["e.code empty"];
                                return
                            }

                            // 設置 成功
                            $scope.inputClass = CSSHasSuccess;
                            Code = val;
                        }
                    }
                ]
            );
        }

        // 運行 angular 模塊
        angular.bootstrap(App, ['app']);
    }
    return function (context) {
        let SystemInfo = context.SystemInfo;
        Lange = context.Lange;
        App = context.App;
        switch (SystemInfo.Register) {
            case Register.Disabled:
                console.log("Register.Disabled")
                break;
            case Register.Open:
                console.log("Register.Open")
                runApp(false);
                break;
            case Register.Invite:
                console.log("Register.Invite")
                runApp(true);
                break;
            default:
                console.error("unknow SystemInfo.Register", SystemInfo.Register)
        }
    };
})
