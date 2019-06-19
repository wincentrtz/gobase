package querybuilder

type queryBuilder struct {
	table string
	where string
	value string
}

type QueryBuilder interface {
	Table(string) QueryBuilder
	Select(string) QueryBuilder
	Where(string, string, string) QueryBuilder
	Excecute() string
}

func NewQueryBuilder() QueryBuilder {
	return &queryBuilder{
		table: "",
		value: "",
	}
}

func (qb *queryBuilder) Table(table string) QueryBuilder {
	qb.table = "FROM " + table
	return qb
}

func (qb *queryBuilder) Select(value string) QueryBuilder {
	qb.value = "SELECT " + value
	return qb
}

func (qb *queryBuilder) Where(attribute string, comparator string, value string) QueryBuilder {
	qb.where = "WHERE " + attribute + " " + comparator + " " + value
	return qb
}

func (qb *queryBuilder) Excecute() string {
	return qb.value + " " + qb.table + " " + qb.where
}
