'use strict';

dcmsControllers.controller('DashboardCtrl', function DashboardCtrl($scope) {
	$scope.sendMessage = function() {
		console.log('Send message...');
		notificationChannel.send('aapje', {field1: 'aapje123', message: "Rofl!"});
	};// empty?
	notificationChannel.on('aapje', function(data) {
		console.log('Recieved an aapje', data);
	});
});