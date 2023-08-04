package main

import (
	"database/sql"
	"github.com/pedrazzani/grpc-product-server/src/database"
	"github.com/pedrazzani/grpc-product-server/src/pb"
	"github.com/pedrazzani/grpc-product-server/src/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "../../resources/db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	productDb := database.NewProduct(db)
	categoryService := services.NewProductService(*productDb)

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
