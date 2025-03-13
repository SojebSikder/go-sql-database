# Description

SQL Database engine created with golang

# Usage

Run queries from file
```
go run main.go run query.sql
```

Run the database cli

```
go run main.go cli
```

Example Queries

```bash
SQL> CREATE TABLE users name age
Table created: users

SQL> INSERT INTO users VALUES Sojeb 24
Row inserted into: users

SQL> INSERT INTO users VALUES Sikder 30
Row inserted into: users

SQL> SELECT FROM users
Data from table: users
[Sojeb 24]
[Sikder 30]
SQL> SELECT FROM users WHERE age 24
Filtered data from table: users
[Sojeb 24]
```
