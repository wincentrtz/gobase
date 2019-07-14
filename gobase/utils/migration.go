package utils

type migrationBuilder struct {
	table   string
	columns []Column
}

type MigrationBuilder interface {
	Table(string) MigrationBuilder
	Column([]Column) MigrationBuilder
	Build() string
}

func Migration() MigrationBuilder {
	return &migrationBuilder{
		table: "",
	}
}

func (mb *migrationBuilder) Table(table string) MigrationBuilder {
	mb.table = table
	return mb
}

func (mb *migrationBuilder) Column(columns []Column) MigrationBuilder {
	mb.columns = columns
	return mb
}

func (mb *migrationBuilder) Build() string {
	var attributes string
	for i, column := range mb.columns {
		if column.name == "" {
			panic("Name Can't be null")
		}
		attributes += column.name + " " + column.typeData + " " + column.constraint
		if i < len(mb.columns)-1 {
			attributes += ","
		}
	}
	return "CREATE TABLE " + mb.table + "(" + attributes + ")"
}
