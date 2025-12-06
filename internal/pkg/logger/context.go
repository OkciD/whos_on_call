package logger

import "context"

type loggerFieldsKey struct{}

func AddFieldsToContext(ctx context.Context, fields Fields) context.Context {
	f := getFieldsFromContext(ctx)

	f.Extend(fields)

	return context.WithValue(ctx, loggerFieldsKey{}, f)
}

func getFieldsFromContext(ctx context.Context) Fields {
	fieldsFromCtx, ok := ctx.Value(loggerFieldsKey{}).(Fields)
	if !ok {
		return Fields{}
	}

	return fieldsFromCtx
}
