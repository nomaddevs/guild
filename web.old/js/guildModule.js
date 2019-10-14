var app = angular.module('guild', ['ngRoute', 'ngCookies']);


app.config(function($routeProvider, $locationProvider) {
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

	$locationProvider.html5Mode(true);
});

app.controller("guildController", ['$scope', '$http', '$cookies', '$location', function($scope, $http, $cookies, $location) {
	var ItoClasses = function(c) {
		switch(c) {
		case 0:
			return "DeathKnight";
		case 1:
			return "DemonHunter";
		case 2:
			return "Druid";
		case 3:
			return "Hunter";
		case 4:
			return "Mage";
		case 5:
			return "Monk";
		case 6:
			return "Paladin";
		case 7:
			return "Priest";
		case 8:
			return "Rogue";
		case 9:
			return "Shaman";
		case 10:
			return "Warlock";
		case 11:
			return "Warrior";
		default:
			return "Unknown";
		}
	};

	var ItoRanks = function(c) {
		switch(c) {
		case 0:
			return "Guild Master";
		case 1:
			return "Raid Leader";
		case 2:
			return "Officer";
		case 3:
			return "Member";
		case 4:
			return "Trial";
		default:
			return "Unknown";
		}
	}

	var ItoRaces = function(c) {
		switch(c) {
		case 2:
			return "Orc";
		case 5:
			return "Undead";
		case 6:
			return "Tauren";
		case 8:
			return "Troll";
		case 9:
			return "Goblin";
		case 10:
			return "Blood Elf";
		case 26:
			return "Pandaren";
		case 27:
			return "Nightborne";
		case 28:
			return "Highmountain Tauren";
		default:
			return "Unknown";
		}
	}

	$scope.Guild = [];
	
	$scope.User = {
		ID: -1,
		BattleTag: "",
		LoggedIn: false,
		Applied: false,
		Characters: [],
	};

	$scope.Application = {};

	$scope.Apply = function() {
		$scope.Application.BattleID = $scope.User.ID;
		$scope.Application.BattleTag = $scope.User.BattleTag;
		console.log($scope.Application);

		$http.post({
	    headers: {
		'Content-Type': 'application/json',
	    },
            method : 'POST',
            url : '/api/beta/apply',
            data : $scope.Application,
        })
        .then(function(response) {
        	$scope.User.Applied = true;
		alert("you did it");
        	console.log(response.data);
        }, function(response) {
        	console.log(response.data);
        });
	}

	$scope.ChangeView = function(view) {
	 	$location.path(view);
	};

	$scope.LoadRoster = function() {
		$http.get("/api/beta/roster")
		.then(function(response) {
			for(var i = 0; i < response.data.members.length; i++) {
				response.data.members[i].character.class = ItoClasses(response.data.members[i].character.class);
				response.data.members[i].character.race = ItoRaces(response.data.members[i].character.race);
				response.data.members[i].rank = ItoRanks(response.data.members[i].rank);
			}
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
    				$scope.User.ID = response.data.ID;
    				$scope.User.BattleTag = response.data.BattleTag;
    				$scope.User.Applied = response.data.Applied;
    				$scope.User.IsAdmin = response.data.Controls == null;
				}, function (response) {
					$scope.User.LoggedIn = false;
				});

				$http.get("/api/beta/apply")
				.then(function(response) {
				$scope.User.Characters = response.data.characters.filter(function(c) {
					return c.level >= 110;
				});
				console.log($scope.User);
				}, function(response) {
					console.log(response);
				});
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
