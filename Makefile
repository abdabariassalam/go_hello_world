build:
	go build -o app

build_docker:
	docker build -t simple_server .

run_docker:
	docker run simple_server

docker_compose:
	docker-compose up -d

test:
	go test
	
run:
	./app