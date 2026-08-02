package db

import (
	"context"
	"database/sql"
)

// интерфейс, которому удовлетворяют sql.DB и sql.Tx
type Executor interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type executorCtxKey struct{}

type WithExecutor struct {
	executor Executor
}

func (we *WithExecutor) GetExecutor(ctx context.Context) Executor {
	if executorFromCtx := getExecutorFromCtx(ctx); executorFromCtx != nil {
		return executorFromCtx
	}

	return we.executor
}

func NewWithExecutor(e Executor) WithExecutor {
	return WithExecutor{
		executor: e,
	}
}

func storeExecutor(ctx context.Context, e Executor) context.Context {
	return context.WithValue(ctx, executorCtxKey{}, e)
}

func getExecutorFromCtx(ctx context.Context) Executor {
	if executorFromCtx, ok := ctx.Value(executorCtxKey{}).(Executor); ok {
		return executorFromCtx
	}

	return nil
}
