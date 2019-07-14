package utils

type Column struct {
	name       string
	typeData   string
	constraint string
}

type columnBuilder struct {
	name       string
	typeData   string
	constraint string
}

type ColumnBuilder interface {
	Name(string) ColumnBuilder
	Primary() ColumnBuilder
	TypeData(string) ColumnBuilder
	NotNull() ColumnBuilder
	Build() Column
}

func NewColumn() ColumnBuilder {
	return &columnBuilder{
		name:       "",
		typeData:   "",
		constraint: "",
	}
}

func (c *columnBuilder) Name(name string) ColumnBuilder {
	c.name = name
	return c
}

func (c *columnBuilder) TypeData(typeData string) ColumnBuilder {
	c.typeData = typeData
	return c
}

func (c *columnBuilder) Primary() ColumnBuilder {
	c.typeData = "serial PRIMARY KEY"
	return c
}

func (c *columnBuilder) NotNull() ColumnBuilder {
	c.constraint = "NOT NULL"
	return c
}

func (c *columnBuilder) Build() Column {
	return Column{
		name:       c.name,
		typeData:   c.typeData,
		constraint: c.constraint,
	}
}
