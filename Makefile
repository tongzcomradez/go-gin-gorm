db.migrate:
	go run cmd/migration/main.go

dev:
	gin --port 3000 --appPort 3030 run main.go