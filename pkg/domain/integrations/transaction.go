package integrations

import (
	"context"
	"database/sql"
)

type ITransaction interface {
	BeginTx(context.Context) (*sql.Tx, error)
	CommitTx(*sql.Tx) error
	RollbackTx(*sql.Tx) error
}
