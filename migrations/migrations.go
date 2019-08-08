package migrations

func GetMigrations() []string {
	return []string{
		UserSchema(),
	}
}
