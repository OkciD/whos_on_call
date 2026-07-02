package apiserver

import (
	"context"

	"github.com/OkciD/whos_on_call/cmd/server/apiserver/gen"
	callStatusDelivery "github.com/OkciD/whos_on_call/internal/server/callstatus/delivery/http"
	deviceDelivery "github.com/OkciD/whos_on_call/internal/server/device/delivery/http"
	deviceFeatureDelivery "github.com/OkciD/whos_on_call/internal/server/devicefeature/delivery/http"
	userDelivery "github.com/OkciD/whos_on_call/internal/server/user/delivery/http"
)

type ApiServer struct {
	UserDelivery          *userDelivery.Handler
	DeviceDelivery        *deviceDelivery.Handler
	DeviceFeatureDelivery *deviceFeatureDelivery.Handler
	CallStatusDelivery    *callStatusDelivery.Handler
}

func (s ApiServer) CreateDevice(ctx context.Context, request gen.CreateDeviceRequestObject) (gen.CreateDeviceResponseObject, error) {
	return s.DeviceDelivery.CreateDevice(ctx, request)
}

func (s ApiServer) UpsertDeviceFeature(ctx context.Context, request gen.UpsertDeviceFeatureRequestObject) (gen.UpsertDeviceFeatureResponseObject, error) {
	return s.DeviceFeatureDelivery.Update(ctx, request)
}

func (s ApiServer) GetStatus(ctx context.Context, request gen.GetStatusRequestObject) (gen.GetStatusResponseObject, error) {
	return s.CallStatusDelivery.GetStatus(ctx, request)
}

func (s ApiServer) GetUser(ctx context.Context, request gen.GetUserRequestObject) (gen.GetUserResponseObject, error) {
	return s.UserDelivery.GetUser(ctx, request)
}
