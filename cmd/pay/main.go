package main

import (
	"fmt"
	"log"
	"net"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/db/postgresql"
	payments "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/payments"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8082")
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

	payments.RegisterPaymentsServiceServer(server, payments.CreatePaymentStore(db.GetDB()))

	fmt.Println("starting server at :8082")
	server.Serve(lis)
}
