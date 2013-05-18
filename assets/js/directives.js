'use strict';

/* Directives */

var app = angular.module('dcms.directives', []);

app.directive('markdown', function () {
    return {
        restrict:'E',
        require: 'ngModel',
        scope: {
            jlValue: '=ngModel'
        },
        template:
            '<div>' +
                '<textarea class="autosize" ng-model="jlValue">'+
                '</textarea>' +
                '<pre ng-show="isEditMode" class="preview" ng-bind-html-unsafe="jlValue | markdown"></pre>' +
            '</div>',
        link: function(scope, elm, attrs) {
            scope.isEditMode = true;
            scope.$watch(attrs['ngModel'], function(){
                $('.autosize',elm).trigger('autosize');
            });
            $('.autosize',elm).autosize({append: "\n"});

            scope.change = function(){
                alert('changed');
            };
            scope.switchToPreview = function () {
                scope.isEditMode = false;
            };
            scope.switchToEdit = function () {
                scope.isEditMode = true;
            };
        }
    };
}).filter('markdown',function() {
        var converter = new Showdown.converter();
        return function(value) {
            return converter.makeHtml(value || '');
        };
    });
