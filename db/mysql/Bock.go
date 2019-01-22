package mysql

type Bock struct {
	table string
}

func (b *Bock) Query(table string) string {
	return "query " + table
}
