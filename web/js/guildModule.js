var app = angular.module('guild', ["ngRoute"]);


app.config(function($routeProvider) {
	$routeProvider
	.when("/", {
		templateUrl: "html/news.html"
	})
	.when("/home", {
		templateUrl: "html/news.html"
	})
	.when("/about", {
		templateUrl: "html/about.html"
	})
	.when("/media", {
		templateUrl: "html/media.html"
	})
	.when("/roster", {
		templateUrl: "html/roster.html"
	})
	.when("/apply", {
		templateUrl: "html/apply.html"
	});
});


app.controller("guildController",['$scope', '$location', function($scope, $location) {
	//$scope.MenuOptions = { "Home": 1, "About": 2, "Media": 3, "Roster": 4, "Apply": 5, };
	//$scope.ActivePage = $scope.MenuOptions.Home;

	$scope.ChangeView = function(view) {
	 	//$scope.ActivePage = view;
	 	$location.path(view);
	};
}]);

app.directive('guildHeader', function(){
	return{
		templateUrl: 'html/header.html',
	};
})
.directive('guildMenuBar', function() {
	return {
		templateUrl: 'html/menu.html',
	};
})
.directive('guildFooter', function() {
	return {
		templateUrl: 'html/footer.html',
	};
})
.directive('guildContent', function() {
	return {
		templateUrl: function(elem, attr) {
			return "html/" + attr.guildPage + ".html";
		}
	};
});
