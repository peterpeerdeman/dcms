'use strict';

dcmsControllers.controller('DocumentTypeOverviewCtrl', function OverviewCtrl($scope, DocumentTypeStorage) {
    $scope.newName = '';
    $scope.documentTypes = DocumentTypeStorage.query();
    $scope.documentType = '';

    $scope.addDocumentType = function () {
        if (!$scope.newName.length) {
            return;
        }

        DocumentTypeStorage.save({id: $scope.newName});
        $scope.newName = '';
        $scope.documentTypes = DocumentTypeStorage.query();
    };
});

dcmsControllers.controller('NewDocumentTypeCtrl', function NewDocumentCtrl($scope, $routeParams, DocumentTypeStorage, $location) {
    $scope.documentType = {"Name": ""};

    $scope.createDocumentType = function() {
        DocumentTypeStorage.save($scope.documentType, function (d) {
            $location.url('/document-type/edit/'  + d.Id);
        });
    }
});

dcmsControllers.controller('EditDocumentTypeCtrl', function EditDocumentCtrl($scope, $routeParams, DocumentTypeStorage, $location, $compile) {
    $scope.documentTypes = DocumentTypeStorage.query();
    $scope.documentType = DocumentTypeStorage.get({id: $routeParams.Id});
    $scope.newDocumentField = {};

    $scope.addField = function() {
        if ($scope.documentType.Fields == null){
            $scope.documentType.Fields = [];
        }
        $scope.documentType.Fields.push($scope.newDocumentField);
        $scope.newDocumentField = {};
    };

    $scope.saveDocumentType = function() {
        $scope.documentType.$update({id: $scope.documentType.Id});
    };
    $scope.deleteDocumentType = function() {
        $scope.documentType.$delete({id: $scope.documentType.Id});
        $location.url('/');
    };
    $scope.removeField = function(field) {
        var index = $scope.documentType.Fields.indexOf(field);
        $scope.documentType.Fields.splice(index,1);
    };
});