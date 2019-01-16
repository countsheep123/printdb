package printdb

import (
	"database/sql"
	"database/sql/driver"
)

func init() {
	sql.Register("printdb", &Driver{})
}

type Driver struct {
	Logger func(name string, args ...interface{})
}

func (d *Driver) Open(name string) (driver.Conn, error) {
	d.Logger("Driver.Open")
	return &Conn{
		Logger: d.Logger,
	}, nil
}

func (d *Driver) OpenConnector(name string) (driver.Connector, error) {
	d.Logger("Driver.OpenConnector")
	return &Connector{
		Logger: d.Logger,
	}, nil
}
