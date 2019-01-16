package printdb

import (
	"database/sql/driver"
	"reflect"
)

type Rows struct {
	Logger func(name string, args ...interface{})
}

func (r *Rows) Columns() []string {
	r.Logger("Rows.Columns")
	return []string{}
}

func (r *Rows) Close() error {
	r.Logger("Rows.Close")
	return nil
}

func (r *Rows) Next(dest []driver.Value) error {
	r.Logger("Rows.Next")
	return nil
}

func (r *Rows) ColumnTypeDatabaseTypeName(index int) string {
	r.Logger("Rows.ColumnTypeDatabaseTypeName")
	return ""
}

func (r *Rows) ColumnTypeLength(index int) (length int64, ok bool) {
	r.Logger("Rows.ColumnTypeLength")
	return 0, true
}

func (r *Rows) ColumnTypeNullable(index int) (nullable, ok bool) {
	r.Logger("Rows.ColumnTypeNullable")
	return true, true
}

func (r *Rows) ColumnTypePrecisionScale(index int) (precision, scale int64, ok bool) {
	r.Logger("Rows.ColumnTypePrecisionScale")
	return 0, 0, true
}

func (r *Rows) ColumnTypeScanType(index int) reflect.Type {
	r.Logger("Rows.ColumnTypeScanType")
	return nil
}

func (r *Rows) HasNextResultSet() bool {
	r.Logger("Rows.HasNextResultSet")
	return false
}

func (r *Rows) NextResultSet() error {
	r.Logger("Rows.NextResultSet")
	return nil
}
