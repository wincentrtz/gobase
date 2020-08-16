package migrations

var migrationList = GetAllMigrations()

func GetMigrations() []string {
	m := make([]string, len(migrationList))
	for _, value := range migrationList {
		m = append(m, value)
	}
	return m
}

func GetMigrationFromSchemaName(schemaName string) []string {
	return []string{migrationList[schemaName]}
}
