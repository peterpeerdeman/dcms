'use strict';

dcmsControllers.controller('DashboardCtrl', function DashboardCtrl($scope, Socket) {
	$scope.sendMessage = function() {
		console.log('Send message...');
		Socket.send('aapje', {field1: 'aapje123', message: "Rofl!"});
	};// empty?
	Socket.on('aapje', function(data) {
		console.log('Recieved an aapje', data);
	});
});