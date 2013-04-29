'use strict';


// Declare app level module which depends on filters, and services
angular.module('dcms', ['dcms.filters', 'dcms.services', 'dcms.directives', 'dcms.controllers']).
  config(['$routeProvider', function($routeProvider) {
    $routeProvider.when('/', {templateUrl: '/assets/angularTemplates/overview.html', controller: 'OverviewCtrl'});
    $routeProvider.when('/:Id', {templateUrl: '/assets/angularTemplates/detail.html', controller: 'DocumentDetailCtrl'});
    $routeProvider.otherwise({redirectTo: '/'});
  }]);
