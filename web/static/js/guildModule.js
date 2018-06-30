var app = angular.module('guild', ["ngRoute"]);

app.config(function($routeProvider) {
	$routeProvider
	.when("/", {
		templateUrl: "news.html"
	})
	.when("/about", {
		templateUrl: "about.html"
	})
	.when("/media", {
		templateUrl: "media.html"
	})
	.when("/roster", {
		templateUrl: "roster.html"
	})
	.when("/apply", {
		templateUrl: "apply.html"
	});
});

app.controller("guildController",['$scope', function($scope) {
	$scope.MenuOptions = { "Home": 1, "About": 2, "Media": 3, "Roster": 4, "Apply": 5, };
	$scope.ActivePage = $scope.MenuOptions.Home;

	$scope.ChangeView = function(activePage) {
	 	$scope.ActivePage = activePage;
	 	console.log($scope.ActivePage);
	};
}]);

app.directive('guildHeader', function(){
	return{
		templateUrl: 'header.html',
	};
})
.directive('guildMenuBar', function() {
	return {
		templateUrl: 'menu.html',
	};
})
.directive('guildFooter', function() {
	return {
		templateUrl: 'footer.html',
	};
})
.directive('guildContent', function() {
	return {
		templateUrl: function(elem, attr) {
			return attr.guildPage + ".html";
		}
	};
});
