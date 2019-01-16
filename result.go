package printdb

type Result struct {
	Logger func(name string, args ...interface{})
}

func (r *Result) LastInsertId() (int64, error) {
	r.Logger("Result.LastInsertId")
	return 0, nil
}

func (r *Result) RowsAffected() (int64, error) {
	r.Logger("Result.RowsAffected")
	return 0, nil
}
