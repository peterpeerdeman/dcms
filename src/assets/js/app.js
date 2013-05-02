'use strict';

angular.module('dcms', ['dcms.filters', 'dcms.services', 'dcms.directives', 'dcms.controllers']).
  config(['$routeProvider', function($routeProvider) {
    $routeProvider.when('/', {templateUrl: '/assets/angularTemplates/dashboard.html', controller: 'DashboardCtrl'});
    $routeProvider.when('/document/overview', {templateUrl: '/assets/angularTemplates/document-overview.html', controller: 'OverviewCtrl'});
    $routeProvider.when('/document/new', {templateUrl: '/assets/angularTemplates/document-new.html', controller: 'NewDocumentCtrl'});
    $routeProvider.when('/document/edit/:Id', {templateUrl: '/assets/angularTemplates/document-edit.html', controller: 'EditDocumentCtrl'});
    $routeProvider.when('/template/overview', {templateUrl: '/assets/angularTemplates/template-overview.html', controller: 'TemplateOverviewCtrl'});
    $routeProvider.when('/pages/overview', {templateUrl: '/assets/angularTemplates/pages-overview.html', controller: 'OverviewCtrl'});
    $routeProvider.when('/sitemap/overview', {templateUrl: '/assets/angularTemplates/sitemap-overview.html', controller: 'OverviewCtrl'});
    $routeProvider.when('/document-type/overview', {templateUrl: '/assets/angularTemplates/document-type-overview.html', controller: 'OverviewCtrl'});
    $routeProvider.when('/channel/overview', {templateUrl: '/assets/angularTemplates/channel-overview.html', controller: 'OverviewCtrl'});
    $routeProvider.otherwise({redirectTo: '/'});
  }]);
