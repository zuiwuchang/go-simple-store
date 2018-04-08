function InitJsPage() {
    var CSSHasNone = "";
    var CSSHasError = "has-error";
    var CSSHasSuccess = "has-success";

    var app = angular.module("app", []);
    // 註冊 數據共享 服務
    app.service("share", function () {
        // 已經 驗證的 合法 可註冊 信息
        this.Email = null;
        this.Password = null;
    });
    app.controller("form",
        [
            "$scope", "$rootScope","$log", "share",
            function ($scope,$rootScope, $log, share) {
                $scope.onSubmit = function () {
                    if (!share.Email) {
                        $log.info("email not set")
                    } else if (!share.Password) {
                        $log.info("Password not set")
                    }

                    // 禁用 ui
                    $rootScope.$broadcast('disabledChange', true);

                    // 啟用 ui
                    $rootScope.$broadcast('disabledChange', false);
                }
            }
        ]
    );
    app.controller("ctrl-email",
        [
            "$scope", "$rootScope", "share",
            function ($scope, $rootScope, share) {
                $scope.disabled = false;
                $scope.$on("disabledChange", function (evt, yes) {
                    $scope.disabled = yes;
                });
                $scope.onChange = function () {
                    // 重置 email 驗證
                    share.Email = null;
                    $scope.inputClass = CSSHasNone;
                };
                $scope.onBlur = function () {
                    // 驗證用戶名
                    var email = $scope.email;
                    

                    // 驗證 註冊
                    $scope.inputClass = CSSHasNone;

                    $scope.inputClass = CSSHasSuccess;
                    // 設置 數據成功
                    share.Email = $scope.email;
                };
            }
        ]
    );
}