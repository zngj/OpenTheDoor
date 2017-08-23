package sqlx

func Query(query string, args ...interface{}) *result {
	return NewDao().Query(query, args...)
}

func Exec(query string, args ...interface{}) error {
	return NewDao().Exec(query, args...)
}
