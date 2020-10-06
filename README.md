# Gobase

Gobase is a boilerplate for Go (golang) for creating boilerplate code, such as **creating table**,
**migration on database**, **creating request and response**, etc.

Currently, _gobase_ is only supported for _postgresql_, but we will add more feature further.

## List of commands

### Serve

serve the applications

```
  go run main.go serve
```

### Generate

Generate boilerplate for code

- #### domain: Create a template for specific domain.

  > **Options:**
  >
  > - domain: domain name

  ```
  go run main.go generate domain [OPTIONS]
  ```

- #### migration: Create a migration for specific domain.

  > **Options:**
  >
  > - domain: domain name

  ```
  go run main.go generate migration [OPTIONS]
  ```

- #### response: Create a custom response.

  > **Options:**
  >
  > - response: response name

  ```
  go run main.go generate response [OPTIONS]
  ```

- #### request: Create a custom request.

  > **Options:**
  >
  > - request: request name

  ```
  go run main.go generate request [OPTIONS]
  ```

### Db

- #### fresh: Drop table on database

  ```
  go run main.go db fresh
  ```

- #### clear: Remove table on database

  ```
  go run main.go db clear
  ```

- #### migrate: Run migration
  ```
  go run main.go db migrate
  ```
