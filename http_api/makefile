## Build all binaries -> command is : make build | make run
build:

	go build -o bin/credentials credentials.go
	go build -o bin/models objects_models.go
	go build -o bin/models ../main.go
	go build -o bin/database_ops database_ops.go
    go build -o bin/models controllers.go


run:
	go run ../main.go

all: build run