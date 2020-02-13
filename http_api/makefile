## Build all binaries -> command is : make build | make run
build:

	go build -o bin/credentials Credentials/credentials.go
	go build -o bin/models Models/models.go
	go build -o bin/models ../main.go
	go build -o bin/database_ops Database/database_ops.go
    go build -o bin/models Routing/routes.go

run:
	go run ../main.go

all: build run