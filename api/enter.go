package api

import (
	echoapi "AuroraPixel/api/echo_api"
	imagesapi "AuroraPixel/api/images_api"
	systemapi "AuroraPixel/api/system_api"
)

type ApiGroup struct {
	SystemApi systemapi.SystemApi
	ImagesApi imagesapi.ImagesApi
	EchoApi   echoapi.EchoApi
}

var GroupApi = new(ApiGroup)
