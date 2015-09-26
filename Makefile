deps:
	go get -u github.com/go-martini/martini
	go get -u github.com/nats-io/nats
dev-deps:
	go get -u github.com/smartystreets/goconvey/convey
build:
	go build 
test:
	go test
