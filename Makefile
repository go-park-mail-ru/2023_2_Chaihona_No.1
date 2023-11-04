test:
	go test -coverpkg=./... -coverprofile=c.out ./.../; go tool cover -func=c.out; go tool cover -html=cov.html
build:
	go build cmd/api/main.go
dc:
	docker-compose up -d
run: dc build
	./main
stop:
	docke-compose down