package printdb

import (
	"context"
	"database/sql/driver"
)

type Stmt struct {
	Logger func(name string, args ...interface{})
}

func (s *Stmt) Close() error {
	s.Logger("Stmt.Close")
	return nil
}

func (s *Stmt) NumInput() int {
	s.Logger("Stmt.NumInput")
	return -1
}

func (s *Stmt) Exec(args []driver.Value) (driver.Result, error) {
	s.Logger("Stmt.Exec", args)
	return &Result{
		Logger: s.Logger,
	}, nil
}

func (s *Stmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	s.Logger("Stmt.ExecContext", args)
	return &Result{
		Logger: s.Logger,
	}, nil
}

func (s *Stmt) Query(args []driver.Value) (driver.Rows, error) {
	s.Logger("Stmt.Query", args)
	return &Rows{
		Logger: s.Logger,
	}, nil
}

func (s *Stmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	s.Logger("Stmt.QueryContext", args)
	return &Rows{
		Logger: s.Logger,
	}, nil
}

func (s *Stmt) ColumnConverter(idx int) driver.ValueConverter {
	s.Logger("Stmt.ColumnConverter")
	return driver.DefaultParameterConverter
}
