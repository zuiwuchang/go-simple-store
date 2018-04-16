define(["angular"], function () {
    'use strict';
    return function (context) {
        let App = context.App;
        let Lange = context.Lange;
        let SystemInfo = context.SystemInfo;

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

        app.controller("ctrl-register",
            [
                "$scope", "$http",
                function ($scope, $http) {
                    let registerMode = SystemInfo.Register.toString();
                    $scope.registerMode = registerMode;
                    $scope.disabled = false;
                    $scope.error = "";
                    $scope.idDisabled = function () {
                        return $scope.disabled || ($scope.registerMode == registerMode);
                    };
                    $scope.hideError = function () {
                        return ($scope.error == "") || ($scope.registerMode == registerMode);
                    };
                    $scope.onSave = function () {
                        // 禁用 ui
                        $scope.disabled = true;
                        $scope.error = "";
                        let doError = (emsg) => {
                            $scope.error = emsg;
                        };

                        // 請求修改
                        $http({
                            method: 'post',
                            url: '/Root/AjaxSaveRegister',
                            data: {
                                registerMode: $scope.registerMode,
                            },
                        }).then((response) => {
                            let rs = response.data;
                            if (rs.Code) {
                                doError(rs.Emsg);
                            } else {
                                registerMode = $scope.registerMode;
                            }
                        }, (response) => {
                            doError(response.data.description)
                        }).finally(() => {
                            // 啟用 ui
                            $scope.disabled = false;
                        });
                    };
                }
            ]
        );
        app.controller("ctrl-smtp",
            [
                "$scope", "$http",
                function ($scope, $http) {
                    let smtp = SystemInfo.SMTP;
                    let email = SystemInfo.Email;
                    let pwd = SystemInfo.Password;

                    $scope.smtp = smtp;
                    $scope.email = email;
                    $scope.pwd = pwd;

                    $scope.disabled = false;
                    $scope.error = "";
                    $scope.testOK = false;
                    $scope.idDisabledTest = function () {
                        return $scope.disabled ||
                            $scope.smtp == "" ||
                            $scope.email == "" ||
                            $scope.pwd == "";
                    }
                    $scope.idDisabledSave = function () {
                        return $scope.disabled ||
                            $scope.smtp == "" ||
                            $scope.email == "" ||
                            $scope.pwd == "" ||
                            ($scope.smtp == smtp && $scope.email == email && $scope.pwd == pwd);
                    };

                    $scope.onTest = function () {
                        // 禁用 ui
                        $scope.disabled = true;
                        $scope.error = "";
                        $scope.testOK = false;

                        // 請求測試
                        $http({
                            method: 'post',
                            url: '/Root/AjaxTestSMTP',
                            data: {
                                smtp: $scope.smtp,
                                email: $scope.email,
                                pwd: $scope.pwd,
                            },
                        }).then((response) => {
                            let rs = response.data;
                            if (rs.Code) {
                                $scope.error = rs.Emsg;
                            } else {
                                $scope.testOK = true;
                            }
                        }, (response) => {
                            $scope.error = response.data.description;
                        }).finally(() => {
                            // 啟用 ui
                            $scope.disabled = false;
                        });
                    };
                    $scope.onSave = function () {
                        // 禁用 ui
                        $scope.disabled = true;
                        $scope.error = "";
                        $scope.testOK = false;

                        // 請求修改
                        $http({
                            method: 'post',
                            url: '/Root/AjaxSaveSMTP',
                            data: {
                                smtp: $scope.smtp,
                                email: $scope.email,
                                pwd: $scope.pwd,
                            },
                        }).then((response) => {
                            let rs = response.data;
                            if (rs.Code) {
                                $scope.error = rs.Emsg;
                            } else {
                                smtp = $scope.smtp;
                                email = $scope.email;
                                pwd = $scope.pwd;
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
        app.controller("ctrl-email",
            [
                "$scope", "$http",
                function ($scope, $http) {
                    let title = SystemInfo.ActiveTitle;
                    let text = SystemInfo.ActiveText;

                    $scope.title = title;
                    $scope.text = text;

                    $scope.disabled = false;
                    $scope.error = "";
                    $scope.testValue = "";

                    $scope.idDisabledTest = function () {
                        return $scope.disabled ||
                            $scope.title == "" ||
                            $scope.text == "";
                    };
                    $scope.idDisabledSave = function () {
                        return $scope.disabled ||
                            $scope.title == "" ||
                            $scope.text == "" ||
                            ($scope.title == title && $scope.text == text);
                    };

                    $scope.onTest = function () {
                        // 禁用 ui
                        $scope.disabled = true;
                        $scope.error = "";
                        $scope.testValue = "";

                        // 請求測試
                        $http({
                            method: 'post',
                            url: '/Root/AjaxTestActive',
                            data: {
                                text: $scope.text,
                            },
                        }).then((response) => {
                            let rs = response.data;
                            if (rs.Code) {
                                $scope.error = rs.Emsg;
                            } else {
                                $scope.testValue = rs.Str;
                            }
                        }, (response) => {
                            $scope.error = response.data.description;
                        }).finally(() => {
                            // 啟用 ui
                            $scope.disabled = false;
                        });
                    };
                    $scope.onSave = function () {
                        // 禁用 ui
                        $scope.disabled = true;
                        $scope.error = "";
                        $scope.testValue = "";

                        // 請求修改
                        $http({
                            method: 'post',
                            url: '/Root/AjaxSaveActive',
                            data: {
                                title: $scope.title,
                                text: $scope.text,
                            },
                        }).then((response) => {
                            let rs = response.data;
                            if (rs.Code) {
                                $scope.error = rs.Emsg;
                            } else {
                                title = $scope.title;
                                text = $scope.text;
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
        // 運行 angular 模塊
        angular.bootstrap(App, ['app']);
    };
});