#!make

CTR_REGISTRY    ?= allenlsy
CTR_TAG         ?= latest

build:
	go build -v -o ./bin/redis-client .
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./bin/redis-client-linux .

docker-build: build
	docker build . -t $(CTR_REGISTRY)/redis-client:$(CTR_TAG)

docker-push: docker-build
	docker push $(CTR_REGISTRY)/redis-client:$(CTR_TAG)
