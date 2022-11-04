run:
	docker-compose up
	
stop:
	docker-compose down

delete:
	docker rmi postgres:latest golang:1.19 alpine:latest onelab4_app:latest


# build:
# 	docker build -t icrudtmpl:multistage -f Dockerfile .
# 	docker image prune

# run:
# 	docker run -d -p 8080:8080 --rm --name ccrudtmpl icrudtmpl:multistage

# stop:
# 	docker stop ccrudtmpl

.DEFAULT_GOAL := run
