function InitJsPage() {
    var app = angular.module("app", []);
    app.controller('form', function ($scope) {
        $scope.email = "123";
        $scope.pwd = 456;
        $scope.onSubmit = function () {
            alert("submit")
        }
        $scope.onBlurEmail = function () {
            alert("blur")
        }
        $scope.onEmailChange = function () {
            console.log($scope.email)
        }
    });
}