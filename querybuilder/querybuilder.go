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
	AndWhere(string, string, string) QueryBuilder
	OrWhere(string, string, string) QueryBuilder
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
	conditionalQuery := attribute + " " + comparator + " '" + value + "'"
	qb.where = "WHERE " + conditionalQuery
	return qb
}

func (qb *queryBuilder) AndWhere(attribute string, comparator string, value string) QueryBuilder {
	conditionalQuery := attribute + " " + comparator + " '" + value + "'"
	qb.where = "AND " + conditionalQuery
	return qb
}

func (qb *queryBuilder) OrWhere(attribute string, comparator string, value string) QueryBuilder {
	conditionalQuery := attribute + " " + comparator + " '" + value + "'"
	qb.where = "OR " + conditionalQuery
	return qb
}

func (qb *queryBuilder) Excecute() string {
	return qb.value + " " + qb.table + " " + qb.where
}
