package main

import (
	"fmt"
	"log"
	"net"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"

	sessrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln("cant listet port", err)
	}

	server := grpc.NewServer()

	sessrep.RegisterAuthCheckerServer(server, sessrep.CreateRedisSessionStorage(sessrep.NewPool(fmt.Sprintf("%s:%s", configs.RedisServerIP, configs.RedisServerPort))))

	fmt.Println("starting server at :8081")
	server.Serve(lis)
}