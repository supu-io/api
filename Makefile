build:
	go build
deps:
	go get -u github.com/go-martini/martini
	go get -u github.com/nats-io/nats
	go get -u github.com/supu-io/messages
dev-deps:
	go get -u github.com/golang/lint/golint
	go get -u github.com/smartystreets/goconvey/convey
test:
	go test
lint:
	golint
cover:
	go test -cover
