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

services.factory('TemplateStorage', function($resource){
    var DEFAULT_ACTIONS = {
        'get':    {method:'GET'},
        'post':   {method:'POST'},
        'update': {method:'PUT'},
        'getAll':  {method:'GET', isArray: true},
        'delete': {method:'DELETE'}
    };
    return $resource('/rest/template/:id', {id: '@id'}, DEFAULT_ACTIONS);
});

services.factory('DocumentTypeStorage', function($resource){
    var DEFAULT_ACTIONS = {
        'get':    {method:'GET'},
        'post':   {method:'POST'},
        'update': {method:'PUT'},
        'getAll':  {method:'GET', isArray: true},
        'delete': {method:'DELETE'}
    };
    return $resource('/rest/document-type/:id', {id: '@id'}, DEFAULT_ACTIONS);
});

services.factory('FileStorage', function($resource){
    var DEFAULT_ACTIONS = {
        'get':    {method:'GET'},
        'post':   {method:'POST'},
        'update': {method:'PUT'},
        'query':  {method:'GET', isArray: true},
        'delete': {method:'DELETE'}
    };
    return $resource('/rest/file/:id', {id: '@id'}, DEFAULT_ACTIONS);
});

services.factory('Socket', function($resource){
    var notificationChannel = $.websocket("ws://localhost:8080/notification", {
        open: function() { console.log('websocket open'); },
        close: function() { console.log('websocket closed'); },
        events: {
            ping: function(e) {
                console.log('got ping', e);
            },
            configuration: function(e) {
                console.log('Configuration was changed', e);
            }
        }
    });
    return notificationChannel;
});