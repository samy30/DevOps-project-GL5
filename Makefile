install:
	cd app; go mod init devopsProjectModule.com/gl5; go mod tidy
clean:
	cd app; rm go.mod go.sum
build: install
	docker-compose build
run:
	docker-compose up
run-detached:
	docker-compose up -d
stop:
	docker-compose down
