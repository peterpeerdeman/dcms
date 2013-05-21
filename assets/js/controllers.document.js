'use strict';

dcmsControllers.controller('DocumentOverviewCtrl', function OverviewCtrl($scope, DocumentStorage) {
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
});

dcmsControllers.controller('NewDocumentCtrl', function NewDocumentCtrl($scope, $routeParams, DocumentStorage, $location, DocumentTypeStorage) {
    $scope.document = {"Name": "", "Type": ""};
    $scope.documentTypes = DocumentTypeStorage.query();

    $scope.createDocument = function() {
        DocumentStorage.save($scope.document, function (d) {
            $location.url('/document/edit/' + d.Id);
        });
    }
});

dcmsControllers.controller('EditDocumentCtrl', function EditDocumentCtrl($scope, $routeParams, DocumentStorage, $location, DocumentTypeStorage) {
    $scope.count = {};
    $scope.documents = DocumentStorage.query();
    $scope.document = DocumentStorage.get({id: $routeParams.Id}, function () {
        $scope.documentType = DocumentTypeStorage.get({id: $scope.document.Type}, function () {
            var fields = $scope.documentType.Fields;
            for (var i = 0; i < fields.length; i++) { 
                var field = fields[i];
                var subfields = [];
                for (var subfield_index = 0; subfield_index < fields[i].Max; subfield_index++) {
                    subfields[subfield_index] = {"index": subfield_index, "required": subfield_index < fields[i].Min};
                }
                $scope.documentType.Fields[i].subfields = subfields;
                if ($scope.document.Fields[field.Name] == undefined) {
                    $scope.document.Fields[field.Name] = new Array();
                } 
            }

            for (var i=0; $scope.documentType.Fields.length > i; i++){
                var field = $scope.documentType.Fields[i];
                $scope.count[field.Name] = field.Min;
            }

        });
    });

    $scope.saveDocument = function() {
        $scope.document.$update({id: $scope.document.Id});
        $location.url('/document/overview');
    };

    $scope.deleteDocument = function() {
        $scope.document.$delete({id: $scope.document.Id});
        $location.url('/document/overview');
    };

    $scope.isVisible = function(name, index){
        return ($scope.document.Fields[name][index] !== undefined && $scope.document.Fields[name][index].length > 0) || $scope.count[name] > index ;
    };

    $scope.isAddable = function(name, max){
        console.log($scope.count[name]);
        return max > $scope.count[name];
    };

    $scope.addField = function(fieldName){
        $scope.count[fieldName]++;
    };

});