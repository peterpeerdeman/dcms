'use strict';

angular.module('dcms', ['dcms.filters', 'dcms.services', 'dcms.directives', 'dcms.controllers']).
  config(['$routeProvider', function($routeProvider) {
    $routeProvider.when('/', {templateUrl: '/cms/assets/angularTemplates/dashboard.html', controller: 'DashboardCtrl'});
    $routeProvider.when('/document/overview', {templateUrl: '/cms/assets/angularTemplates/document-overview.html', controller: 'OverviewCtrl'});
    $routeProvider.when('/document/new', {templateUrl: '/cms/assets/angularTemplates/document-new.html', controller: 'NewDocumentCtrl'});
    $routeProvider.when('/document/edit/:Id', {templateUrl: '/cms/assets/angularTemplates/document-edit.html', controller: 'EditDocumentCtrl'});
    $routeProvider.when('/template/overview', {templateUrl: '/cms/assets/angularTemplates/template-overview.html', controller: 'TemplateOverviewCtrl'});
    $routeProvider.when('/pages/overview', {templateUrl: '/cms/assets/angularTemplates/pages-overview.html', controller: 'OverviewCtrl'});
    $routeProvider.when('/sitemap/overview', {templateUrl: '/cms/assets/angularTemplates/sitemap-overview.html', controller: 'OverviewCtrl'});
    $routeProvider.when('/document-type/overview', {templateUrl: '/cms/assets/angularTemplates/document-type-overview.html', controller: 'OverviewCtrl'});
    $routeProvider.when('/channel/overview', {templateUrl: '/cms/assets/angularTemplates/channel-overview.html', controller: 'OverviewCtrl'});
    $routeProvider.otherwise({redirectTo: '/'});
  }]);
