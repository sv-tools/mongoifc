BREWFILE=./.github/Brewfile

NO_COLOR=\033[0m
OK_COLOR=\033[32;01m

DOCKER_IMAGE="mongo"
DOCKER_NAME="mongoifc-test"
MONGO_PORT=27888
MONGO_USERNAME="admin"
MONGO_PASSWORD="adminpass"
MONGO_URI="mongodb://$(MONGO_USERNAME):$(MONGO_PASSWORD)@localhost:$(MONGO_PORT)/?authSource=admin"

ifeq ($(shell uname), Darwin)
all: brew-install
endif

all: go-install generate-mocks tidy lint test done

done:
	@echo "$(OK_COLOR)==> Done.$(NO_COLOR)"

brew-install:
	@echo "$(OK_COLOR)==> Checking and installing dependencies using brew...$(NO_COLOR)"
	@brew bundle --file $(BREWFILE)

go-install:
	@echo "$(OK_COLOR)==> Checking and installing dependencies using go install...$(NO_COLOR)"
	@go install github.com/golang/mock/mockgen@v1

run-test:
	@echo "$(OK_COLOR)==> Testing...$(NO_COLOR)"
	@MONGO_URI=$(MONGO_URI) richgo test -cover -race ./...

stop-docker:
	@echo "$(OK_COLOR)==> Stopping docker...$(NO_COLOR)"
	@docker rm --force $(DOCKER_NAME) || true

run-docker:
	@echo "$(OK_COLOR)==> Running docker...$(NO_COLOR)"
	@docker rm --force $(DOCKER_NAME) || true
	@docker run -d --name=$(DOCKER_NAME) -p $(MONGO_PORT):27017 -e MONGO_INITDB_ROOT_USERNAME=$(MONGO_USERNAME) -e MONGO_INITDB_ROOT_PASSWORD=$(MONGO_PASSWORD) $(DOCKER_IMAGE)

test: run-docker run-test stop-docker

lint:
	@echo "$(OK_COLOR)==> Linting via golangci-lint...$(NO_COLOR)"
	@golangci-lint run --fix ./...

tidy:
	@echo "$(OK_COLOR)==> Updating go.mod...$(NO_COLOR)"
	@go mod tidy

run-mockgen:
	@mockgen -destination=mocks/gomock/mocks.go -package mocks . ChangeStream,Client,Collection,Cursor,Database,IndexView,Session,SingleResult,SessionContext

run-mockery:
	@mockery --all --output mocks/mockery --disable-version-string --case underscore

generate-mocks: run-mockgen run-mockery
