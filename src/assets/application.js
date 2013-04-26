var m = angular.module('dcms', ['dcmsServices']);
m.config(function($interpolateProvider) {
  $interpolateProvider.startSymbol('[[');
  $interpolateProvider.endSymbol(']]');
});

var services = angular.module('dcmsServices', ['ngResource']);
services.factory('Query', function($resource) {
  var DEFAULT_ACTIONS = {
        'get':    {method:'GET'},
        'save':   {method:'POST'},
        'update': {method:'PUT'},
        'query':  {method:'GET', isArray: true},
        'remove': {method:'DELETE'},
        'delete': {method:'DELETE'}
      };
  return $resource('/rest/query/:id', {id: '@id'}, DEFAULT_ACTIONS);
});

function QueryCtrl($scope, $http, Query) {
  $scope.showEdit = 0;
  $scope.master = {};
  $scope.queries = Query.query();
  $scope.queryOrder = 'Id';
  $scope.edit = function(query) {
    $scope.master = query;
    $scope.current = angular.copy(query);
    $scope.showEdit = 1;
  }
  $scope.add = function() {
      $scope.current = {Id: ''};
      $scope.showEdit = 1;
    }
  $scope.delete = function(query) {
    for (var i = 0; i < $scope.queries.length; i++) {
        if ($scope.queries[i].Id === query.Id) {
           $scope.queries.splice(i, 1);
           i--;
        }
    }
    query.$remove({id: query.Id});
  }
  $scope.save = function() {
    if ($scope.current.Id != '') {
        $scope.current.$update({id: $scope.current.Id});
        for (var i = 0; i < $scope.queries.length; i++) {
            if ($scope.queries[i].Id === $scope.current.Id) {
               $scope.queries[i] = $scope.current;
            }
        }
    } else {
        var query = Query.save($scope.current);
        $scope.queries.push(query);
    }
    $scope.showEdit = 0;
  }
  $scope.cancel = function() {
    $scope.master = {};
    $scope.showEdit = 0;
  };
}