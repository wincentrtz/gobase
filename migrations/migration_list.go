package migrations

func GetAllMigrations() map[string]string {
	return map[string]string{
		"user": UserSchema(),
	}
}
