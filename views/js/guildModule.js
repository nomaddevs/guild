var app = angular.module('guild', []);

app.controller("guildController", function($scope) {
	$scope.MenuOptions = ["home", "media", "roster", "apply"];
	$scope.User = {
		BattleTag: 'Guest',
		LoggedIn: false
	};
});

app.directive('guildHeader', function(){
	return{
		templateUrl: 'header.html',
	};
})
.directive('guildMenuBar', function() {
	return {
		templateUrl: 'navbar.html',
	};
})
.directive('guildContent', function() {
	return {
		templateUrl: function(elem, attr) {
			return attr.guildPage + ".html";
		}
	};
})
.directive('guildFooter', function() {
	return {
		templateUrl: 'footer.html',
	};
});
