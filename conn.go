package printdb

import (
	"context"
	"database/sql/driver"
)

type Conn struct {
	Logger func(name string, args ...interface{})
}

func (c *Conn) Ping(ctx context.Context) error {
	c.Logger("Conn.Ping")
	return nil
}

func (c *Conn) Prepare(query string) (driver.Stmt, error) {
	c.Logger("Conn.Prepare", query)
	return &Stmt{
		Logger: c.Logger,
	}, nil
}

func (c *Conn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	c.Logger("Conn.PrepareContext", query)
	return &Stmt{
		Logger: c.Logger,
	}, nil
}

func (c *Conn) Close() error {
	c.Logger("Conn.Close")
	return nil
}

func (c *Conn) Begin() (driver.Tx, error) {
	c.Logger("Conn.Begin")
	return &Tx{
		Logger: c.Logger,
	}, nil
}

func (c *Conn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	c.Logger("Conn.Begin")
	return &Tx{
		Logger: c.Logger,
	}, nil
}

func (c *Conn) Exec(query string, args []driver.Value) (driver.Result, error) {
	c.Logger("Conn.Exec", query, args)
	return &Result{
		Logger: c.Logger,
	}, nil
}

func (c *Conn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	c.Logger("Conn.ExecContext", query, args)
	return &Result{
		Logger: c.Logger,
	}, nil
}

func (c *Conn) Query(query string, args []driver.Value) (driver.Rows, error) {
	c.Logger("Conn.Query", query, args)
	return &Rows{
		Logger: c.Logger,
	}, nil
}

func (c *Conn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	c.Logger("Conn.QueryContext", query, args)
	return &Rows{
		Logger: c.Logger,
	}, nil
}

func (c *Conn) ResetSession(ctx context.Context) error {
	c.Logger("Conn.ResetSession")
	return nil
}
