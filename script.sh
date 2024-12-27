#go install -tags ‘mysql,posgres,mongodb’ github.com/golang-migrate/migrate/v4/cmd/migrate@latest

#create migrate
migrate create -ext sql -dir db/migrations create_table_first

migrate -database "mysql://root:korie123@tcp(mysql-db:3306)/krs_management" -path db/migrations up

migrate -database "mysql://root:korie123@tcp(mysql-db:3306)/krs_management" -path db/migrations down


docker compose up -d

docker exec -it krs-management-go-app-1 bash

go test -v -run TestDataInsertion1 -timeout 30m #process
go test -v -run TestDataInsertion2 -timeout 30m #
go test -v -run TestDataInsertion3 -timeout 30m

