package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type TxManager interface {
	WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

type txManager struct {
	db *sql.DB
}

func (txm txManager) WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	if alreadyExistingExecutor := getExecutorFromCtx(ctx); alreadyExistingExecutor != nil {
		return fn(ctx)
	}

	tx, err := txm.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
	}()

	ctxWithExecutor := storeExecutor(ctx, tx)

	err = fn(ctxWithExecutor)
	if err != nil {
		rbError := tx.Rollback()
		if rbError != nil {
			return fmt.Errorf("error rollbacking a transaction: %w", errors.Join(err, rbError))
		}

		return fmt.Errorf("error executing fn in transaction: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error commiting transaction: %w", err)
	}

	return nil
}

func NewTxManager(db *sql.DB) TxManager {
	return txManager{
		db: db,
	}
}
