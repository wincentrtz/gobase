package migrations

func Schema() string {
	return `CREATE TABLE users(
		id serial PRIMARY KEY,
		name VARCHAR NOT NULL,
		email VARCHAR NOT NULL,
		created_on TIMESTAMP NOT NULL
	);`
}
