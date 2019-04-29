BINARY=server
test: 
	go test -v -cover -covermode=atomic ./...
dev-run: 
	 DATABASE_CONNECTION_STRING='postgres://postgres:@localhost:32768/comments_groot?true' go run main.go
vendor:
	@dep ensure -v

server: vendor
	go build -o ${BINARY}

install: 
	go build -o ${BINARY}

unittest:
	go test -short $$(go list ./... | grep -v /vendor/)

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t go-clean-arch .

run:
	docker-compose up -d

stop:
	docker-compose down

.PHONY: clean install unittest build docker run stop vendor