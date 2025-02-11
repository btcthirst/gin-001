exe_path := bin/app
work_path := cmd/app/main.go

.PHONY: build
build:
	go build -o ./${exe_path} ./${work_path}

.PHONY: run
run: build
	./${exe_path}

.PHONY: up
up:
	docker-compose up --build

.PHONY: down
down:
	docker-compose down --rmi all