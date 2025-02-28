NO_COLOR=\033[0m
OK_COLOR=\033[32;01m

DOCKER_IMAGE="mongoifc"
DOCKER_FILE=".github/test.dockerfile"
DOCKER_NAME="mongoifc-test"
MONGO_PORT=27888
MONGO_USERNAME="admin"
MONGO_PASSWORD="adminpass"
MONGO_URI="mongodb://$(MONGO_USERNAME):$(MONGO_PASSWORD)@127.0.0.1:$(MONGO_PORT)/?authSource=admin&directConnection=true"

run-test:
	@echo "$(OK_COLOR)==> Testing...$(NO_COLOR)"
	@MONGO_URI=$(MONGO_URI) go test -cover -race ./...

stop-docker:
	@echo "$(OK_COLOR)==> Stopping docker...$(NO_COLOR)"
	@docker rm --force $(DOCKER_NAME) || true

docker-build:
	@docker rm --force $(DOCKER_NAME) || true
	@docker build . --rm --file $(DOCKER_FILE) --tag $(DOCKER_IMAGE)

run-docker: docker-build
	@echo "$(OK_COLOR)==> Running docker...$(NO_COLOR)"
	@docker run -d --name=$(DOCKER_NAME) -p $(MONGO_PORT):27017 -e MONGO_INITDB_ROOT_USERNAME=$(MONGO_USERNAME) -e MONGO_INITDB_ROOT_PASSWORD=$(MONGO_PASSWORD) $(DOCKER_IMAGE)
	@sleep $${SLEEP:-4}
	@docker exec $(DOCKER_NAME) /usr/bin/mongosh -u $(MONGO_USERNAME) -p $(MONGO_PASSWORD) --eval "rs.initiate()"

test: run-docker run-test stop-docker
