'use strict';

/* Controllers */

angular.module('dcms.controllers', [])

    .controller('DashboardCtrl', function DashboardCtrl($scope) {
        // empty?
    })

    .controller('DocumentOverviewCtrl', function OverviewCtrl($scope, DocumentStorage) {

        $scope.newName = '';
        $scope.documents = DocumentStorage.query();
        $scope.documentType = '';
        $scope.orderProp = 'Name';

        $scope.addDocument = function () {
            if (!$scope.newName.length) {
                return;
            }

            DocumentStorage.save({id: $scope.newName});
            $scope.newName = '';
            $scope.documents = DocumentStorage.query();
        };
    })

    .controller('NewDocumentCtrl', function NewDocumentCtrl($scope, $routeParams, DocumentStorage, $location, DocumentTypeStorage) {

        $scope.document = {"Name": "", "Type": ""};
        $scope.documentTypes = DocumentTypeStorage.query();

        $scope.createDocument = function() {
            DocumentStorage.save($scope.document, function (d) {
                $location.url('/document/edit/' + d.Id);
            });
        }
    })

    .controller('EditDocumentCtrl', function EditDocumentCtrl($scope, $routeParams, DocumentStorage, $location, DocumentTypeStorage) {
        $scope.documents = DocumentStorage.query();
        $scope.document = DocumentStorage.get({id: $routeParams.Id}, function () {
            $scope.documentType = DocumentTypeStorage.get({id: $scope.document.Type});
        });

        $scope.editDocument = function() {
            $scope.document.$update({id: $scope.document.Id});
            $location.url('/');
        };

        $scope.deleteDocument = function() {
            $scope.document.$delete({id: $scope.document.Id});
            $location.url('/');
        }
    })

    .controller('TemplateOverviewCtrl', function TemplateOverviewCtrl($scope, TemplateStorage){
        $scope.test = 'test123';
        $scope.templates = TemplateStorage.get();
    })

    .controller('DocumentTypeOverviewCtrl', function OverviewCtrl($scope, DocumentTypeStorage) {
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
    })

    .controller('NewDocumentTypeCtrl', function NewDocumentCtrl($scope, $routeParams, DocumentTypeStorage, $location) {
        $scope.documentType = {"Name": ""};

        $scope.createDocumentType = function() {
            DocumentTypeStorage.save($scope.documentType, function (d) {
                $location.url('/document-type/edit/'  + d.Id);
            });
        }
    })

    .controller('EditDocumentTypeCtrl', function EditDocumentCtrl($scope, $routeParams, DocumentTypeStorage, $location, $compile) {
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
