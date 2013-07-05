'use strict';

dcmsControllers.controller('DashboardCtrl', function DashboardCtrl($scope, Socket) {
	$scope.sendMessage = function() {
		Socket.send('aapje', {field1: 'aapje123', message: "Rofl!"});
	};

	Socket.on('aapje', function(data) {
		console.log('Recieved an aapje', data);
	});
});