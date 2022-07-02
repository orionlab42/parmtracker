all: test build run
production: test build
list:
	@grep '^[^#[:space:]].*:' Makefile
update-deps:
	go mod tidy
	go mod vendor
test:
	go test ./testutil
	go test -p 1 -cover ./...
build:
	go install github.com/orionlab42/parmtracker
	cd server/client && npm install && npm run-script build && cd ../..
run:
	parmtracker
nohup:
	nohup parmtracker &
kill:
	killall -9 parmtracker
client:
	cd server/client && npm start
