package repositories

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// UnitOfWork implements UnitOfWork using pgx
type UnitOfWork struct {
	tx pgx.Tx
}

// NewUnitOfWork initializes a new UnitOfWork with the given pool
func NewUnitOfWork(tx pgx.Tx) *UnitOfWork {
	return &UnitOfWork{tx: tx}
}

// Commit commits the transaction
func (u *UnitOfWork) Commit(ctx context.Context) error {
	if u.tx == nil {
		return errors.New("transaction not started")
	}

	err := u.tx.Commit(ctx)
	u.tx = nil
	if err != nil {
		return err
	}
	return nil
}

// Rollback rolls back the transaction
func (u *UnitOfWork) Rollback(ctx context.Context) error {
	if u.tx == nil {
		return errors.New("transaction not started")
	}

	err := u.tx.Rollback(ctx)
	u.tx = nil
	if err != nil && err != pgx.ErrTxClosed {
		return err
	}
	return nil
}

// UnitOfWorkFactory is responsible for creating UnitOfWork instances
type UnitOfWorkFactory struct {
	pool *pgxpool.Pool
}

// NewUnitOfWorkFactory creates a new UnitOfWorkFactory
func NewUnitOfWorkFactory(pool *pgxpool.Pool) *UnitOfWorkFactory {
	return &UnitOfWorkFactory{pool: pool}
}

// Create creates a new UnitOfWork instance
func (f *UnitOfWorkFactory) Create(ctx context.Context) (*UnitOfWork, error) {
	log.Println("create a new unit of work")
	tx, err := f.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return NewUnitOfWork(tx), nil
}
