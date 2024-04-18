package db

import (
	"context"
	"database/sql"
	"go-practice/pkg/domain/integrations"
)

type TransactionHandler struct {
	db *sql.DB
}

func NewTransactionHandler(db *sql.DB) *TransactionHandler {
	return &TransactionHandler{db: db}
}

var _ integrations.ITransaction = (*TransactionHandler)(nil)

// BeginTx implements integrations.ITransaction.
func (t *TransactionHandler) BeginTx(ctx context.Context) (*sql.Tx, error) {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// CommitTx implements integrations.ITransaction.
func (t *TransactionHandler) CommitTx(tx *sql.Tx) error {
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// RollbackTx implements integrations.ITransaction.
func (t *TransactionHandler) RollbackTx(tx *sql.Tx) error {
	err := tx.Rollback()
	if err != nil {
		return err
	}
	return nil
}
