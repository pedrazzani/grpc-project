package services

import (
	"context"
	"github.com/pedrazzani/grpc-product-server/src/database"
	"github.com/pedrazzani/grpc-product-server/src/pb"
)

type ProductService struct {
	pb.UnimplementedProductServiceServer
	ProductDB database.Product
}

func NewProductService(ProductDB database.Product) *ProductService {
	return &ProductService{
		ProductDB: ProductDB,
	}
}

func (p *ProductService) CreateProduct(ctx context.Context, request *pb.CreateProductRequest) (*pb.ProductResponse, error) {
	Product, err := p.ProductDB.Create(request.Name, request.Description)
	if err != nil {
		return nil, err
	}

	ProductResponse := &pb.ProductResponse{
		Id:          Product.ID,
		Name:        Product.Name,
		Description: Product.Description,
	}

	return ProductResponse, nil
}

func (p *ProductService) ListProducts(ctx context.Context, in *pb.Blank) (*pb.ProductList, error) {
	Products, err := p.ProductDB.FindAll()
	if err != nil {
		return nil, err
	}

	var ProductResponseList []*pb.ProductResponse

	for _, Product := range Products {
		ProductRsponse := &pb.ProductResponse{
			Id:          Product.ID,
			Name:        Product.Name,
			Description: Product.Description,
		}

		ProductResponseList = append(ProductResponseList, ProductRsponse)
	}

	return &pb.ProductList{Products: ProductResponseList}, nil
}

func (p *ProductService) FindAllProductsByCategoryId(ctx context.Context, in *pb.CategoryProductRequest) (*pb.ProductList, error) {
	Products, err := p.ProductDB.FindByCategoryId(in.CategoryId)
	if err != nil {
		return nil, err
	}
	var ProductResponseList []*pb.ProductResponse

	for _, Product := range Products {
		ProductRsponse := &pb.ProductResponse{
			Id:          Product.ID,
			Name:        Product.Name,
			Description: Product.Description,
		}

		ProductResponseList = append(ProductResponseList, ProductRsponse)
	}

	return &pb.ProductList{Products: ProductResponseList}, nil
}
