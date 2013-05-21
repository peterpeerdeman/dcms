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

        $scope.editDocument = function() {
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
    })

    .controller('FileuploadCtrl', function FileuploadCtrl($scope, FileStorage) {

        $scope.files = FileStorage.query();

        $scope.uploadFile = function() {
            $scope.uploadStarted = true;
            var formData = new FormData($('form')[0]);
            $.ajax({
                url: '/rest/file',
                type: 'POST',
                xhr: function() {  // custom xhr
                    var myXhr = $.ajaxSettings.xhr();
                    if(myXhr.upload){ // check if upload property exists
                        myXhr.upload.addEventListener('progress',progressHandlingFunction, false); // for handling the progress of the upload
                    }
                    return myXhr;
                },
                //Ajax events
//                beforeSend: beforeSendHandler,
                success: function(){$scope.files = FileStorage.query();},
//                error: errorHandler,
                // Form data
                data: formData,
                //Options to tell JQuery not to process data or worry about content-type
                cache: false,
                contentType: false,
                processData: false
            });

            function progressHandlingFunction(e){
                if(e.lengthComputable){
                    $('progress').attr({value:e.loaded,max:e.total});
                }
            }

        }


    });
