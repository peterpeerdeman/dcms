'use strict';

/* Directives */

var app = angular.module('dcms.directives', []);
app.directive('documentFields', function() {
    return {
        restrict:'C',
        template:
            '<input id="name" placeholder="Title" ng-model="document.Name">'+
            '<input id="documentId" ng-model="document.Id" type="hidden">'
    };
});

app.directive('string', function () {
    return {
        restrict: 'E',
        template: '<input type="text">'
    };
});
