# A CRUD Application with Go and PostgreSQL

https://www.youtube.com/playlist?list=PLASVJZN9c0VXA4ZPib37apjY3ZoHgBe65

## Prerequisite
PostgreSQL v10+ 

Go 1.18+

Running database server:
```
$ sudo service postgresql start
```
For working directly with database, use either pgAdmin 4 or via CLI (demo):
```
$ sudo -i -u postgres
$ pqsl
``` 

$\textbf{Note}$ - The application's demo is ran in the following environment:

```
$ psql --version
psql (PostgreSQL) 12.16 (Ubuntu 12.16-0ubuntu0.20.04.1)

$ go version
go version go1.20.1 linux/amd64
```

Database and its owner has been set up. The default variables below are used for database connection:
```go
var (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "postgres"
	DBNAME   = "postgres"
	PASSWORD = "postgres"
)
```

Demo table is pre-populated with the following schema and data:
```sql
CREATE TABLE IF NOT EXISTS students (
		student_id SERIAL PRIMARY KEY,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		enrollment_date DATE
);

INSERT INTO students (first_name, last_name, email, enrollment_date) VALUES
('John', 'Doe', 'john.doe@example.com', '2023-09-01'),
('Jane', 'Smith', 'jane.smith@example.com', '2023-09-01'),
('Jim', 'Beam', 'jim.beam@example.com', '2023-09-02');
```

## Demo
```
$ charles@CHRLS:/mnt/c/Users/phucn/OneDrive - Carleton University/Documents/CU/F23/COMP3005/A4$ ls
README.md  backend.go  go.mod  go.sum  main.go  schema.go  tui.go

$ go run .
```

### GET
Get all items of the table

<img src="get-demo.gif" width="650" />

### CREATE
Add new item to the table

<img src="create-demo.gif" width="650" />

### UPDATE
Change an item's field

<img src="update-demo.gif" width="650" />

### DELETE
Delete an item from the table

<img src="delete-demo.gif" width="650" />

## File Structure

```
postgres-demo/
---|
   ----- main.go	// driver code
   |
   ----- tui.go		// terminal ui
   |
   ----- backend.go	// contains database connection setup & SQL queries
```

The CRUD application is written in pure Go.  
Connection and queries to PostgreSQL is made possible thanks to Go postgres database driver ```github.com/lib/pq``` as well as Go standard library for generic interface around SQL databases ```database/sql```.  
The client is terminal UI (a text-based interface yet is so much more than raw CLI and, 
when done right, can be more aesthetically pleasing than an average GUI), using external library ```https://github.com/charmbracelet/bubbletea```.