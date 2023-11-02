test:
	go test -coverpkg=./... -coverprofile=c.out ./.../; go tool cover -func=c.out