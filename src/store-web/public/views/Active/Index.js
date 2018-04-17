define(["angular"], function () {
    'use strict';
    return function (context) {
        let App = context.App;

        const interval = context.interval;
        // 創建 angular 模塊
        let app = angular.module('app', []);


        app.controller("ctrl",
            [
                "$scope", "$interval", "$http",
                function ($scope, $interval, $http) {
                    $scope.wait = context.wait;
                    $scope.error = "";
                    $scope.sendOk = false;
                    $scope.disabled = false;
                    $scope.isDisabled = function () {
                        return $scope.disabled || $scope.wait != 0;
                    };

                    let stop;
                    function startInterval() {
                        stop = $interval(function () {
                            if ($scope.wait > 0) {
                                $scope.wait--;
                            }
                            if ($scope.wait == 0) {
                                $interval.cancel(stop);
                            }
                        }, 1000);
                    }
                    if ($scope.wait > 0) {
                        startInterval();
                    }
                    $scope.onSend = function () {
                        // 禁用 ui
                        $scope.disabled = true;
                        $scope.error = "";
                        $scope.sendOk = false;

                        // 發送請求
                        $http({
                            method: 'get',
                            url: '/Active/AjaxResend',
                        }).then((response) => {
                            let rs = response.data;
                            if (rs.Code) {
                                $scope.error = rs.Emsg;
                            } else{
                                // 成功
                                $scope.sendOk = true;
                            }
                        }, (response) => {
                            $scope.error = response.data.description;
                        }).finally(() => {
                            // 啟用 ui
                            $scope.disabled = false;
                            $scope.wait = interval;
                            startInterval();
                        });
                    };
                },
            ]
        );

        // 運行 angular 模塊
        angular.bootstrap(App, ['app']);
    };
});