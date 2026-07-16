package apiserver

import (
	callStatusDelivery "github.com/OkciD/whos_on_call/internal/server/callstatus/delivery/http"
	deviceDelivery "github.com/OkciD/whos_on_call/internal/server/device/delivery/http"
	deviceFeatureDelivery "github.com/OkciD/whos_on_call/internal/server/devicefeature/delivery/http"
	userDelivery "github.com/OkciD/whos_on_call/internal/server/user/delivery/http"
)

type ApiServer struct {
	userDelivery.UserHandler
	deviceDelivery.DeviceHandler
	deviceFeatureDelivery.DeviceFeatureHandler
	callStatusDelivery.CallStatusHandler
}
