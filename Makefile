test:
	go test -coverpkg=./... -coverprofile=c.out ./.../; go tool cover -func=c.out; go tool cover -html=cov.html
build:
	go build -o ./bin/api cmd/api/main.go;
	go build -o ./bin/auth cmd/auth/main.go;
	go build -o ./bin/posts cmd/posts/main.go;
	go build -o ./bin/pay cmd/pay/main.go;
docker_build: 
	docker build -t m0rdovorot/kopilka.api -f API.Dockerfile .;
	docker build -t m0rdovorot/kopilka.auth -f auth.Dockerfile .;
	docker build -t m0rdovorot/kopilka.posts -f posts.Dockerfile .;
	docker build -t m0rdovorot/kopilka.pay -f pay.Dockerfile .;
docker_run:
	docker run -d -p 8001:8001 kopilka.api;
	docker run -d -p 8001:8081 kopilka.auth;
	docker run -d -p 8001:8082 kopilka.pay;
	docker run -d -p 8001:8083 kopilka.posts;
run:
	docker-compose up -d
stop:
	docker-compose down
