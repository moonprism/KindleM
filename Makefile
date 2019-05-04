build:
	swag init
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o ./docker/app/run .
	cp config.yml docker/app/
	sudo docker-compose build
	rm docker/app/run docker/app/config.yml
run:
	sudo docker-compose up -d

restart:
	sudo docker-compose down
	sudo docker-compose up -d