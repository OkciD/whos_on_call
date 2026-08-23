package context

import "context"

type requestIDCtxKey = struct{}

func StoreRequestID(ctx context.Context, reqID string) context.Context {
	return context.WithValue(ctx, requestIDCtxKey{}, reqID)
}

func GetRequestID(ctx context.Context) string {
	reqID, ok := ctx.Value(requestIDCtxKey{}).(string)
	if !ok {
		return "undef"
	}

	return reqID
}
