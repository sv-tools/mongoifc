BREWFILE=./.github/Brewfile

NO_COLOR=\033[0m
OK_COLOR=\033[32;01m

DOCKER_IMAGE="mongoifc"
DOCKER_FILE=".github/test.dockerfile"
DOCKER_NAME="mongoifc-test"
MONGO_PORT=27888
MONGO_USERNAME="admin"
MONGO_PASSWORD="adminpass"
MONGO_URI="mongodb://$(MONGO_USERNAME):$(MONGO_PASSWORD)@127.0.0.1:$(MONGO_PORT)/?authSource=admin&directConnection=true"

all: go-install generate-mocks tidy lint test done

done:
	@echo "$(OK_COLOR)==> Done.$(NO_COLOR)"

go-install:
	@echo "$(OK_COLOR)==> Checking and installing dependencies using go install...$(NO_COLOR)"
	@go install go.uber.org/mock/mockgen@latest
	@go install github.com/vektra/mockery/v2@v2
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/kyoh86/richgo@latest

run-test:
	@echo "$(OK_COLOR)==> Testing...$(NO_COLOR)"
	@MONGO_URI=$(MONGO_URI) richgo test -cover -race ./...

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

lint:
	@echo "$(OK_COLOR)==> Linting via golangci-lint...$(NO_COLOR)"
	@golangci-lint run --fix ./...

tidy:
	@echo "$(OK_COLOR)==> Updating go.mod...$(NO_COLOR)"
	@go mod tidy -compat=1.21

run-mockgen:
	@mockgen -destination=mocks/gomock/mocks.go -package mocks . ChangeStream,Client,Collection,Cursor,Database,IndexView,Session,SingleResult,SessionContext,ClientEncryption

run-mockery:
	@mockery --all --with-expecter --srcpkg github.com/sv-tools/mongoifc --output mocks/mockery --disable-version-string --case underscore

clean-mocks:
	@rm -rf mocks

generate-mocks: clean-mocks run-mockgen run-mockery
