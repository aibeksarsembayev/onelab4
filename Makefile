run-nodocker:
	go run ./cmd

build:
	docker build -t icrudtmpl:multistage -f Dockerfile .
	docker image prune

run:
	docker run -d -p 8080:9090 --rm --name ccrudtmpl icrudtmpl:multistage

stop:
	docker stop ccrudtmpl

.DEFAULT_GOAL := build
