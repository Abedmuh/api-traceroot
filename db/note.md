# command
migrate -database "postgres://postgres:pass@localhost:5432/abhvps?sslmode=disable" -path db/migrations down
migrate -database "postgres://postgres:pass@localhost:5432/abhvps?sslmode=disable" -path db/migrations up
migrate create -ext sql -dir db/migrations -seq create_serverList_tables