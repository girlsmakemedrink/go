DB_DSN := "postgres://postgres:1234@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

gen:
	oapi-codegen -config openapi/.openapi -include-tags messages -package messages openapi/openapi.yaml > ./internal/web/messages/api.gen.go

# Таргет для создания новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations -seq -digits 1 ${NAME}

# Применение миграций
migrate:
	$(MIGRATE) -verbose up

# Откат миграций
migrate-down:
	$(MIGRATE) down

# для удобства добавим команду run, которая будет запускать наше приложение
run:
	go run cmd/app/main.go

lint:
	golangci-lint run --out-format=colored-line-number