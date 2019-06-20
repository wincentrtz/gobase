package querybuilder

import "fmt"

type queryBuilder struct {
	table string
	where string
	join  string
	value string
}

type QueryBuilder interface {
	Table(string) QueryBuilder
	Select(string) QueryBuilder
	Where(string, string, string) QueryBuilder
	AndWhere(string, string, string) QueryBuilder
	OrWhere(string, string, string) QueryBuilder
	Join(string, string, string) QueryBuilder
	Build() string
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
	qb.where = qb.where + " AND " + conditionalQuery
	return qb
}

func (qb *queryBuilder) OrWhere(attribute string, comparator string, value string) QueryBuilder {
	conditionalQuery := attribute + " " + comparator + " '" + value + "'"
	qb.where = "OR " + conditionalQuery
	return qb
}

func (qb *queryBuilder) Join(table string, firstKey string, secondKey string) QueryBuilder {
	qb.join = "JOIN " + table + " ON " + firstKey + " = " + secondKey
	return qb
}

func (qb *queryBuilder) Build() string {
	fmt.Println(qb.value + " " + qb.table + " " + qb.join + " " + qb.where)
	return qb.value + " " + qb.table + " " + qb.join + " " + qb.where
}
