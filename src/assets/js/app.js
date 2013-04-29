'use strict';


// Declare app level module which depends on filters, and services
angular.module('dcms', ['dcms.filters', 'dcms.services', 'dcms.directives', 'dcms.controllers']).
  config(['$routeProvider', function($routeProvider) {
    $routeProvider.when('/', {templateUrl: '/assets/angularTemplates/overview.html', controller: 'OverviewCtrl'});
    $routeProvider.when('/document/:Id', {templateUrl: '/assets/angularTemplates/detail.html', controller: 'DocumentDetailCtrl'});
    $routeProvider.when('/templateEditor/:Id', {templateUrl: '/assets/angularTemplates/templateEditor.html', controller: 'TemplateEditorCtrl'});
    $routeProvider.when('/templateEditor', {templateUrl: '/assets/angularTemplates/templateOverview.html', controller: 'TemplateOverviewCtrl'});
    $routeProvider.otherwise({redirectTo: '/'});
  }]);
