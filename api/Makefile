APP_NAME := api

CWD := $(shell pwd)

.PHONY: init
init:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/zeromicro/go-zero/tools/goctl@latest
	go install github.com/wqyjh/goctl-gogen@v0.1.1
	npm install -g npm
	npm install -g swagger2openapi

.PHONY: api
api:
	goctl api plugin -plugin goctl-gogen="--home ${CWD}/../template" -api def/main.api -dir ./ -style goZero

.PHONY: doc
doc:
	swag fmt
	swag init -g api.go
	swagger2openapi -o docs/swagger.json docs/swagger.json

.PHONY: model-cached
model-cached:
	goctl model mysql ddl -src="model/sql/model_cached.sql" -dir="./model" -c -style goZero

.PHONY: model
model: model-cached
	goctl model mysql ddl -src="model/sql/model.sql" -dir="./model" -style goZero

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s -X main.version=`git rev-parse --short HEAD`" -o ${APP_NAME}

.PHONY: clean
clean:
	@if [ -f ${APP_NAME} ] ; then rm ${APP_NAME} ; fi
