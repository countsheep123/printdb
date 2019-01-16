package printdb

type Tx struct {
	Logger func(name string, args ...interface{})
}

func (t *Tx) Commit() error {
	t.Logger("Tx.Commit")
	return nil
}

func (t *Tx) Rollback() error {
	t.Logger("Tx.Rollback")
	return nil
}
