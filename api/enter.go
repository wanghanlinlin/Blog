package api

import systemapi "AuroraPixel/api/system_api"

type ApiGroup struct {
	SystemApi systemapi.SystemApi
}

var GroupApi = new(ApiGroup)
