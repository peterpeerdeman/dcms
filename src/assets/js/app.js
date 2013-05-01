'use strict';


// Declare app level module which depends on filters, and services
angular.module('dcms', ['dcms.filters', 'dcms.services', 'dcms.directives', 'dcms.controllers']).
  config(['$routeProvider', function($routeProvider) {
    $routeProvider.when('/', {templateUrl: '/assets/angularTemplates/overview.html', controller: 'OverviewCtrl'});
    $routeProvider.when('/document/new', {templateUrl: '/assets/angularTemplates/detail-new.html', controller: 'NewDocumentCtrl'});
    $routeProvider.when('/document/edit/:Id', {templateUrl: '/assets/angularTemplates/detail-edit.html', controller: 'EditDocumentCtrl'});
    $routeProvider.when('/templateEditor', {templateUrl: '/assets/angularTemplates/templateEditor.html', controller: 'TemplateEditorCtrl'});
    $routeProvider.otherwise({redirectTo: '/'});
  }]);
