package service

import (
	"context"
	"ms-seller/model"
	pb "ms-seller/proto"
	"ms-seller/repository"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductService struct {
	ProductRepository repository.ProductRepository
	pb.UnimplementedProductServiceServer
}

func NewProductService(ProductRepository repository.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: ProductRepository,
	}
}

func (h *ProductService) AddProduct(ctx context.Context, in *pb.AddProductRequest) (*pb.ProductResponse, error) {
	var input = model.Product{
		SellerID:    int(in.SellerId),
		Name:        in.Name,
		Price:       in.Price,
		Stock:       int(in.Stock),
		Category_id: int(in.CategoryId),
		Discount:    int(in.Discount),
	}

	product, err := h.ProductRepository.Create(&input)
	if err != nil {
		return nil, err
	}

	var response = &pb.ProductResponse{
		Productid:  int32(product.ID),
		SellerId:   int32(product.SellerID),
		Price:      float32(product.Price),
		Stock:      int32(product.Stock),
		CategoryId: int32(product.Category_id),
		Discount:   int32(product.Discount),
	}

	return response, nil
}

func (h *ProductService) GetProducts(ctx context.Context, in *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	Products, err := h.ProductRepository.ReadAll(int(in.SellerId))
	if err != nil {
		return nil, err
	}

	var list []*pb.ProductResponse

	for _, v := range Products {
		u := pb.ProductResponse{
			Productid:  int32(v.ID),
			SellerId:   int32(v.SellerID),
			Price:      float32(v.Price),
			Stock:      int32(v.Stock),
			CategoryId: int32(v.Category_id),
			Discount:   int32(v.Discount),
		}
		list = append(list, &u)
	}

	return &pb.GetProductsResponse{
		Products: list,
	}, nil
}

func (t *ProductService) GetProductByID(ctx context.Context, in *pb.GetProductByIDRequest) (*pb.ProductResponse, error) {
	product, err := t.ProductRepository.ReadID(int(in.Productid), int(in.SellerId))
	if err != nil {
		return nil, err
	}

	var response = &pb.ProductResponse{
		Productid:  int32(product.ID),
		SellerId:   int32(product.SellerID),
		Price:      float32(product.Price),
		Stock:      int32(product.Stock),
		CategoryId: int32(product.Category_id),
		Discount:   int32(product.Discount),
	}

	return response, nil
}

func (t *ProductService) DeleteProduct(ctx context.Context, in *pb.DeleteProductRequest) (*emptypb.Empty, error) {
	err := t.ProductRepository.Delete(int(in.Productid), int(in.SellerId))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (t *ProductService) UpdateProduct(ctx context.Context, in *pb.UpdateProductRequest) (*pb.ProductResponse, error) {
	var product = model.Product{
		SellerID:    int(in.SellerId),
		Name:        in.Name,
		Price:       in.Price,
		Stock:       int(in.Stock),
		Category_id: int(in.CategoryId),
		Discount:    int(in.Discount),
	}

	err := t.ProductRepository.Update(&product)
	if err != nil {
		return nil, err
	}

	var response = &pb.ProductResponse{
		Productid:  int32(product.ID),
		SellerId:   int32(product.SellerID),
		Price:      float32(product.Price),
		Stock:      int32(product.Stock),
		CategoryId: int32(product.Category_id),
		Discount:   int32(product.Discount),
	}

	return response, nil
}
