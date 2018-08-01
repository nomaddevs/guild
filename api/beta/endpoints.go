package beta

var (
	EndpointLogin    = "/login"
	EndpointCallback = "/callback"

	epAPI        = "/api/"
	epVersion    = "beta/"
	EndpointRoot = epAPI + epVersion // 	/api/beta/

	EndpointNews        = EndpointRoot + "news"        // 	/api/beta/news
	EndpointRealms      = EndpointRoot + "realms"      // 	/api/beta/realms
	EndpointAbout       = EndpointRoot + "about"       // 	/api/beta/about
	EndpointMedia       = EndpointRoot + "media"       //	/api/beta/media
	EndpointRoster      = EndpointRoot + "roster"      //   /api/beta/roster
	EndpointApply       = EndpointRoot + "apply"       // 	/api/beta/apply
	EndpointRecruitment = EndpointRoot + "recruitment" //	/api/beta/recruitment
	EndpointProgression = EndpointRoot + "progression" //	/api/beta/progression
	EndpointMythicPlus  = EndpointRoot + "mythicplus"  // 	/api/beta/mythicplus
)
