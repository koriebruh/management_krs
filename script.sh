#go install -tags ‘mysql,posgres,mongodb’ github.com/golang-migrate/migrate/v4/cmd/migrate@latest

#create migrate
migrate create -ext sql -dir db/migrations create_table_first

migrate -database "mysql://root:korie123@tcp(mysql-db:3306)/krs_management" -path db/migrations up

migrate -database "mysql://root:korie123@tcp(mysql-db:3306)/krs_management" -path db/migrations down