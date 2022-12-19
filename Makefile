include .env

default: build

dev: build
	@echo Run application
	@./scripts/run.sh

build: clean
	@echo -n Building...
	@cd "src" || exit 1 && go build -o ../bin/app
	@echo OK
	@echo -n Set permissions to binary...
	@cd "bin" || exit 1 && chmod -x app && chmod 775 app
	@echo OK

install:
	@echo -n Install dependencies...
	@go mod download
	@echo OK

clean:
	@echo -n Clean...
	@rm -f bin/app
	@echo OK
