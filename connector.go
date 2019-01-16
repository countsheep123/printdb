package printdb

import (
	"context"
	"database/sql/driver"
)

type Connector struct {
	Logger func(name string, args ...interface{})
}

func (c *Connector) Connect(ctx context.Context) (driver.Conn, error) {
	c.Logger("Connector.Connect")
	return &Conn{
		Logger: c.Logger,
	}, nil
}

func (c *Connector) Driver() driver.Driver {
	c.Logger("Connector.Driver")
	return &Driver{
		Logger: c.Logger,
	}
}
