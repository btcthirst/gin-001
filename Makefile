exe_path := bin/app
work_path := cmd/app/main.go

.PHONY: build
build:
	go build -o ./${exe_path} ./${work_path}

.PHONY: run
run: build
	./${exe_path}
