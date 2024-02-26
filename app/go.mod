module hello-k8s

go 1.22

replace hello-k8s/api => ./api

replace hello-k8s/app => ./app

replace hello-k8s/data => ./data

replace hello-k8s/env => ./env

require (
	github.com/gorilla/handlers v1.4.0
	github.com/gorilla/mux v1.6.2
	github.com/jinzhu/gorm v1.9.1
)

require (
	github.com/denisenkom/go-mssqldb v0.10.0 // indirect
	github.com/erikstmartin/go-testdb v0.0.0-20160219214506-8d10e4a1bae5 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gorilla/context v1.1.1 // indirect
	github.com/jinzhu/inflection v0.0.0-20180308033659-04140366298a // indirect
	github.com/jinzhu/now v1.1.2 // indirect
	github.com/lib/pq v1.0.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.7 // indirect
)
