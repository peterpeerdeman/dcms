'use strict';

/* Controllers */

angular.module('dcms.controllers', [])

    .controller('OverviewCtrl', function OverviewCtrl($scope, DocumentStorage) {

        $scope.newName = '';
        $scope.documents = DocumentStorage.query();
        $scope.documentType = '';

        $scope.addDocument = function () {
            if (!$scope.newName.length) {
                return;
            }

            DocumentStorage.save({id: $scope.newName});
            $scope.newName = '';
            $scope.documents = DocumentStorage.query();
        };
    })

    .controller('DocumentDetailCtrl', function DocumentDetailCtrl($scope, $routeParams, DocumentStorage, $location) {

        $scope.nieuwsTemplate = [{'name':'title', 'type':'Text'},{'name':'subtitle', 'type':'Text'}];
        $scope.persoonTemplate = [{'name':'title', 'type':'Text'},{'name':'age', 'type':'Text'}];


        $scope.documents = DocumentStorage.query();
        $scope.document = DocumentStorage.get({id: $routeParams.Id});

        $scope.editDocument = function() {
            $scope.document.$update({id: $scope.document.Id});
            $location.url('/');
        };

        $scope.deleteDocument = function() {
            $scope.document.$delete({id: $scope.document.Id});
            $location.url('/');
        }
    })

    .controller('TemplateOverviewCtrl', function TemplateOverviewCtrl($scope, TemplateStorage, $routeParams){
        $scope.templates = TemplateStorage.getAll();
        $scope.template = TemplateStorage.get({id: $scope.document.Id});
        $scope.fields = TemplateStorage.getAll();
    })

.controller('TemplateEditorCtrl', function TemplateEditorCtrl($scope, TemplateStorage, $routeParams){
        $scope.templates = TemplateStorage.getAll();
        $scope.template = TemplateStorage.get({id: $scope.document.Id});
        $scope.fields = TemplateStorage.getAll();
    });
