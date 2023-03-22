package api

import (
	imagesapi "AuroraPixel/api/images_api"
	systemapi "AuroraPixel/api/system_api"
)

type ApiGroup struct {
	SystemApi systemapi.SystemApi
	ImagesApi imagesapi.ImagesApi
}

var GroupApi = new(ApiGroup)
