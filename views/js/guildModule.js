var app = angular.module('guild', []);

app.controller("guildController", function($scope) {
	$scope.MenuOptions = ["Home", "Media", "Roster", "Apply"];
});

app.directive('guildMenuBar', function(){
	return{
		template: '<b>You did it!</b>',
	};
});
