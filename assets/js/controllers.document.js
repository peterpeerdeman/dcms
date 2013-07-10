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

dcmsControllers.controller('EditDocumentCtrl', function EditDocumentCtrl($scope, $routeParams, DocumentStorage, $location, DocumentTypeStorage, FileStorage, Socket) {
    $scope.count = {};
    $scope.documents = DocumentStorage.query();
    var updating = false;
    $scope.document = DocumentStorage.get({id: $routeParams.Id}, function () {
        $scope.documentType = DocumentTypeStorage.get({id: $scope.document.Type}, function () {
            var fields = $scope.documentType.Fields;
            for (var i = 0; i < fields.length; i++) { 
                var field = fields[i];

                // initialize the Field on the document
                if ($scope.document.Fields[field.Name] === undefined) {
                    $scope.document.Fields[field.Name] = new Array();
                } 

                // try and determine the minimum shown subfields
                var min = 0;
                for (var subfield_index = 0; subfield_index < field.Max; subfield_index++) {
                    if (subfield_index < field.Min) {
                        min = subfield_index;
                    }
                    if ($scope.document.Fields[field.Name][subfield_index]) {
                        min = subfield_index;
                    }
                }
                $scope.count[field.Name] = min;

                // contruct the subfields
                var subfields = [];
                for (var subfield_index = 0; subfield_index < field.Max; subfield_index++) {
                    subfields[subfield_index] = {"index": subfield_index, "required": subfield_index <= min};
                }
                $scope.documentType.Fields[i].subfields = subfields;
            }
            console.log($scope.count);

            // watch for changes
            $scope.$watch('document', watchForChanges, true);
            $scope.original = JSON.stringify($scope.document);
            Socket.on('document.'+$scope.document.Id, onRecieveChange);
        });
    });

    $scope.documentDirty = false;
    $scope.documentUpdating = false;

    function watchForChanges(newValue, oldValue) {
        if (newValue && !$scope.documentUpdating) {
            $scope.documentDirty = true;
        }
    }

    function sendChanges(newValue, oldValue) {
        if (newValue) {
            var oldHash = Hashcode.value($scope.original);
            var newJson = JSON.stringify($scope.document);
            var newHash = Hashcode.value(newJson);
            if (oldHash != newHash) {
                var differ = new diff_match_patch();
                var patch = differ.patch_make($scope.original, newJson);
                Socket.send('document.' + $scope.document.Id, {'hash': oldHash, 'patch': patch});
            }
        }
    }

    function onRecieveChange(message) {

        //save caret position
        var focusedElement = document.activeElement;
        var caretPos = focusedElement.selectionStart;

        var prevContent = focusedElement.value;
        var prevContentLength = focusedElement.value.length;

        console.log('Recieved a document update', message);
        $scope.$apply(function () {
            $scope.documentUpdating = true;
            var current = JSON.stringify($scope.document);
            var differ = new diff_match_patch();
            var patchedDocument = differ.patch_apply(message.data.patch, current);
            $scope.document = JSON.parse(patchedDocument[0]);
            $scope.original = JSON.stringify($scope.document);
            $scope.documentUpdating = false;
        });

        var currentContent = focusedElement.value;
        var lenghtChangedSize = currentContent.length - prevContentLength;

        //set caret position
        if (prevContent.slice(0,caretPos) === currentContent.slice(0,caretPos)){
            focusedElement.setSelectionRange(caretPos,caretPos);
        } else {
            focusedElement.setSelectionRange(caretPos+lenghtChangedSize,caretPos+lenghtChangedSize);
        }
    }

    function waiter() {
        if ($scope.documentDirty && !$scope.documentUpdating) {
            $scope.documentDirty = false;
            sendChanges($scope.document, $scope.original);
            $scope.$apply(function () {
                $scope.original = JSON.stringify($scope.document);
            });
        }
        setTimeout(waiter, 1000);
    }
    setTimeout(waiter, 1000);

    $scope.saveDocument = function() {
        var timestamp = new Date();
        var hours = timestamp.getHours();
        var minutes = timestamp.getMinutes();
        var seconds = timestamp.getSeconds();
        var formattedTime = hours + ':' + minutes + ':' + seconds;

        $scope.document.$update({id: $scope.document.Id});
        Socket.send('documentsaved', {documentId: $scope.document.Id, name: $scope.document.Name, timestamp:formattedTime});
        $location.url('/document/overview');
    };

    $scope.deleteDocument = function() {
        $scope.document.$delete({id: $scope.document.Id});
        $location.url('/document/overview');
    };

    $scope.isVisible = function(name, index){
        return $scope.count[name] >= index;
    };

    $scope.isAddable = function(name, max){
        return $scope.count[name] < max - 1;
    };

    $scope.addField = function(fieldName){
        $scope.count[fieldName]++;
    };

    $scope.openAssetpicker = function(){
        if(!$scope.files){
            $scope.files = FileStorage.query();
        }
    };

    $scope.addAssets = function(fieldName){
        $('.asset:checked').each(function(){
            $scope.document.Fields[fieldName].push(this.value);
        });
        $('#assetpicker').modal('hide');
    };

});