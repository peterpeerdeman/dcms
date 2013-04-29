'use strict';

/* Controllers */

angular.module('dcms.controllers', [])

    .controller('OverviewCtrl', function OverviewCtrl($scope, DocumentStorage) {

        $scope.newName = '';
        var documents = $scope.documents = DocumentStorage.query();

        $scope.addDocument = function () {
            if (!$scope.newName.length) {
                return;
            }

            DocumentStorage.save({
                name: $scope.newName
            });
            $scope.newName = '';
            var documents = $scope.documents = DocumentStorage.query();
        };
    })

    .controller('DocumentDetailCtrl', function DocumentDetailCtrl($scope, $routeParams, DocumentStorage) {

        $scope.documents = DocumentStorage.query();
        $scope.document = DocumentStorage.get({id: $routeParams.Id});

        $scope.editDocument = function(){
            $scope.document.$update({id: $scope.document.Id});
        };

        $scope.deleteDocument = function(){
            $scope.document.$delete({id: $scope.document.Id});
        }
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