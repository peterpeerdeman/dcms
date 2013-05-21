'use strict';

dcmsControllers.controller('TemplateOverviewCtrl', function TemplateOverviewCtrl($scope, TemplateStorage){
    $scope.test = 'test123';
    $scope.templates = TemplateStorage.get();
});