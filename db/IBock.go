package db

type IBock interface {
	Query(table string) string
}
