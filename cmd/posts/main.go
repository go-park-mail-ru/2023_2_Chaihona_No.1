package main

import (
	"fmt"
	"log"
	"net"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/db/postgresql"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/posts"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatalln("cant listet port", err)
	}

	server := grpc.NewServer()

	var db postgresql.Database
	err = db.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	posts.RegisterPostsServiceServer(server, posts.CreatePostStore(db.GetDB()))

	fmt.Println("starting server at :8083")
	server.Serve(lis)
}
