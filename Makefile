BINARY=server
test: 
	go test -v -cover -covermode=atomic ./...
dev-run: 
	 DATABASE_CONNECTION_STRING='postgres://postgres:@localhost:32768/comments_groot?true' go run main.go
vendor:
	@dep ensure -v

server: vendor
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /${BINARY} .

install: 
	GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -gcflags "all=-N -l" -o ${BINARY}

unittest:
	go test -short $$(go list ./... | grep -v /vendor/)

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t groot-rest-api .

run:
	docker-compose up --build

stop:
	docker-compose down

.PHONY: clean install unittest build docker run stop vendor