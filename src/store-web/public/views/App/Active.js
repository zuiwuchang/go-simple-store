define(["angular"], function () {
    'use strict';
    return function (context) {
        let App = context.App;

        // 創建 angular 模塊
        let app = angular.module('app', []);


        app.controller("ctrl",
            [
                "$scope", "$interval",
                function ($scope, $interval) {
                    $scope.count = 5;

                    let stop;
                    stop = $interval(function () {
                        if($scope.count > 0){
                            $scope.count--;
                        }
                        if($scope.count == 0){
                            $interval.cancel(stop);

                            window.location.href="/";
                        }
                    }, 1000);
                },
            ]
        );

        // 運行 angular 模塊
        angular.bootstrap(App, ['app']);
    };
});