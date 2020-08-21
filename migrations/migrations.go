package migrations


func GetMigrations() []string {
	m := make([]string, len(MigrationList))
	for _, value := range MigrationList {
		m = append(m, value)
	}
	return m
}

func GetMigrationFromSchemaName(schemaName string) []string {
	return []string{MigrationList[schemaName]}
}
