var app = angular.module('guild', ['ngRoute', 'ngCookies']);


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

app.controller("guildController", ['$scope', '$http', '$cookies', '$location', function($scope, $http, $cookies, $location) {
	$scope.User = {
		BattleTag: "",
		LoggedIn: false,
		Applied: false,
		Characters: [],
	};

	$scope.Guild = [];

	$scope.ChangeView = function(view) {
	 	$location.path(view);
	};

	$scope.LoadRoster = function() {
		$http.get("/api/beta/roster")
		.then(function(response) {
			$scope.Guild = response.data;
			console.log($scope.Guild);
		}, function (response) {
			console.log(response.data);
		});
	};

	var login = function() {
		if(!$scope.User.LoggedIn) {
			var token = $cookies.get("token");
			if(!angular.isUndefined(token)) {
				$http.get("/api/beta/user")
				.then(function(response) {
    				$scope.User.LoggedIn = true;
    				$scope.User.ID = response.data.id;
    				$scope.User.BattleTag = response.data.battletag;
    				$scope.User.Applied = response.data.applied;
    				$scope.User.Characters = response.data.characters.filter(function(c) {
					return c.level >= 110;
				});
				}, function (response) {
					$scope.User.LoggedIn = false;
				});
/*
				$http.get("/api/beta/apply")
				.then(function(response) {
				$scope.User.Characters = response.data.characters.filter(function(c) {
					return c.level >= 110;
				});
				console.log(response.data);
				}, function(response) {
					console.log(response);
				});
*/
			}
		}
	};
	$scope.LoadRoster();
	login();
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
