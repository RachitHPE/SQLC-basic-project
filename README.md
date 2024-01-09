# SQLC-basic-test

SQLC is a database tool to generate idiomatic type-safe go code that derived from our own SQL Queries.

It is basically an alternative for GORM in golang

Install packages-->
1) sqlc
 
   curl -LO https://downloads.sqlc.dev/sqlc_1.25.0_linux_amd64.tar.gz
   
   tar xzf sqlc_1.25.0_linux_amd64.tar.gz
   
   ./sqlc version
   
   /home/rachit/sqlc version

3) migrate
 
   curl -LO https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-arm64.tar.gz
   
   tar xzf migrate.linux-arm64.tar.gz
   
   ./migrate -version
   
   /home/rachit/migrate version

Commands to run-->
1) /home/rachit/sqlc generate --> this creates a folder named db which contains 3 files db.go, models.go and query.sql.go
 
2) /home/rachit/migrate create -ext sql -dir misc/migrations CreateAuthor  --> this create a folder named misc/migrations which has sql.up and sql.down schema file
 
3) /home/rachit/migrate -source "file://misc/migrations" -database "postgres://dbuser:rachit@localhost:5432/testdb?sslmode=disable" up --> this creates a table in the postgres database
 
4) go run main.go --> this starts the golang application which inserts data in the table
