# Gobase

Gobase is a boilerplate for Go (golang) for creating boilerplate code, such as **creating table**, 
**migration on database**, **creating request and response**, etc.

Currently, *gobase* is only supported for _postgresql_, but we will add more feature further. 

## List of commands
### Generate
Generate boilerplate for code
* domain: Create a template for specific domain.

    >**Options:** 
    >- domain: domain name 
    
    ```
    go run main.go generate domain [OPTIONS] 
    ```
* ####migration: Create a migration for specific domain.

    >**Options:** 
    >- domain: domain name 
    
    ```
    go run main.go generate migration [OPTIONS] 
    ```
* ####response: Create a custom response.

    >**Options:** 
    >- response: response name 
    
    ```
    go run main.go generate response [OPTIONS] 
    ```
* ####request: Create a custom request.

     >**Options:** 
     >- request: request name 
    
    ```
    go run main.go generate request [OPTIONS] 
    ```

### Db
* ####fresh: Drop table on database

     >**Options:** 
     >- domain: domain name 
     >- all: all table
        
     ```
     go run main.go db fresh [OPTIONS] 
     ```

* ####clear: Remove table on database
    
     >**Options:** 
     >- domain: domain name (NOTE: domain name in plural, see table name or migration file name)
     >- all: all table
     
     ```
     go run main.go db clear [OPTIONS] 
     ```

* ####migrate: Run migration

     >**Options:** 
     >- migration: migration name (NOTE: migration name in plural, see table name or migration file name)
     >- all: all table
       
     ```
     go run main.go db migrate [OPTIONS] 
     ```






