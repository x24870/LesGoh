DOCKER_IMAGE_APP_NAME=demo_app

build_app:
	@echo "start to build docker image - $(DOCKER_IMAGE_APP_NAME)"
	docker build -t $(DOCKER_IMAGE_APP_NAME) -f docker/Dockerfile .

deploy_app:
	docker run --name demo_app --rm --network demo_app_default -p 8080:8080 --env-file docker/config/.env $(DOCKER_IMAGE_APP_NAME) --restart=always

deploy_db:
	docker run --name postgres --rm -e POSTGRES_USER=test -e POSTGRES_HOST_AUTH_METHOD=trust -e POSTGRES_DB=test -p 5432:5432 -d postgres:14.4-alpine 

up:
	docker-compose up -d

clean: 
	docker ps -aq | xargs docker rm -fv
