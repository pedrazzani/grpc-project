package main

import (
	"database/sql"
	"github.com/pedrazzani/grpc-category-server/src/database"
	"github.com/pedrazzani/grpc-category-server/src/pb"
	"github.com/pedrazzani/grpc-category-server/src/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "../../resources/db/db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := services.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
