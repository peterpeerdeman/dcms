'use strict';

/* Services */

var services = angular.module('dcms.services', ['ngResource']);
services.factory('DocumentStorage', function($resource) {
    var DEFAULT_ACTIONS = {
        'get':    {method:'GET'},
        'save':   {method:'POST'},
        'update': {method:'PUT'},
        'query':  {method:'GET', isArray: true},
        'delete': {method:'DELETE'}
    };
    return $resource('/rest/document/:id', {id: '@id'}, DEFAULT_ACTIONS);
});