build:
	swag init
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o ./docker/app/run .
	cp config.yml docker/app/
	docker-compose build
	rm docker/app/run docker/app/config.yml
run:
	docker-compose up -d

restart:
	docker-compose down
	docker-compose up -d