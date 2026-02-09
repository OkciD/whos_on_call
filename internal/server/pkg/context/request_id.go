package context

import "context"

type requestIdCtxKey = struct{}

func StoreRequestId(ctx context.Context, reqId string) context.Context {
	return context.WithValue(ctx, requestIdCtxKey{}, reqId)
}

func GetRequestId(ctx context.Context) string {
	reqId, ok := ctx.Value(requestIdCtxKey{}).(string)
	if !ok {
		return "undef"
	}

	return reqId
}
