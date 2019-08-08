package db

func Fresh() {
	Drop()
	Migrate()
}
