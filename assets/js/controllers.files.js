'use strict';

dcmsControllers.controller('FileuploadCtrl', function FileuploadCtrl($scope, FileStorage) {

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