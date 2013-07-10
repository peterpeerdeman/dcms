'use strict';

dcmsControllers.controller('DashboardCtrl', function DashboardCtrl($scope, Socket) {
    $scope.messages = [];

	$scope.sendMessage = function() {
		Socket.send('aapje', {field1: 'aapje123', message: "Rofl!"});
	};

	Socket.on('aapje', function(data) {
		console.log('Recieved an aapje', data);
	});
    Socket.on('documentsaved', function(data) {
        $scope.$apply(function(){
            $scope.messages.push({timestamp: data.data.timestamp, documentId: data.data.documentId, documentName:data.data.name});
        });
	});
});