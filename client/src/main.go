package main

import (
	"context"
	"fmt"
	"github.com/pedrazzani/grpc-client/src/pb"
	"google.golang.org/grpc"
)

func main() {

	err, categoryClient := getCategoryClient()
	if err != nil {
		panic(err)
	}

	err, productClient := getProductClient()
	if err != nil {
		panic(err)
	}

	categories, err := categoryClient.ListCategories(context.Background(), &pb.Blank{})
	if err != nil {
		panic(err)
	}
	for _, category := range categories.Categories {
		fmt.Println("Category:", category)

		products, _ := productClient.FindAllProductsByCategoryId(context.Background(), &pb.CategoryProductRequest{CategoryId: category.Id})
		for _, product := range products.Products {
			fmt.Println("  Product:", product)
		}
	}
}

func getCategoryClient() (error, pb.CategoryServiceClient) {
	connection, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewCategoryServiceClient(connection)
	return err, client
}

func getProductClient() (error, pb.ProductServiceClient) {
	connection, err := grpc.Dial("0.0.0.0:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewProductServiceClient(connection)
	return err, client
}
