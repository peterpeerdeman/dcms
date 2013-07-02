'use strict';

dcmsControllers.controller('DashboardCtrl', function DashboardCtrl($scope) {
	$scope.sendMessage = function() {
		console.log('Send message...');
		ws.send('ping', {field1: 'aapje', message: "Rofl!"});
	};// empty?
});