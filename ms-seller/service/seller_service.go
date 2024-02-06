package service

import (
	"context"
	"ms-seller/model"
	"ms-seller/pb"
	"ms-seller/repository"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

type SellerService struct {
	SellerRepository repository.SellerRepository
	pb.UnimplementedSellerServiceServer
}

func NewSellerService(seller repository.SellerRepository) *SellerService {
	return &SellerService{
		SellerRepository: seller,
	}
}

// products service
func (se *SellerService) AddProduct(ctx context.Context, in *pb.AddProductRequest) (*pb.ProductResponse, error) {
	var input = model.Product{
		SellerID:    int(in.SellerId),
		Name:        in.Name,
		Price:       in.Price,
		Stock:       int(in.Stock),
		Category_id: int(in.CategoryId),
		Discount:    int(in.Discount),
	}

	product, err := se.SellerRepository.CreateProduct(&input)
	if err != nil {
		return nil, err
	}

	var response = &pb.ProductResponse{
		Productid:  int32(product.ID),
		SellerId:   int32(product.SellerID),
		Name:       product.Name,
		Price:      float32(product.Price),
		Stock:      int32(product.Stock),
		CategoryId: int32(product.Category_id),
		Discount:   int32(product.Discount),
	}

	return response, nil
}

func (se *SellerService) GetProductsBySeller(ctx context.Context, in *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	products, err := se.SellerRepository.ReadAllProducts(int(in.SellerId))
	if err != nil {
		return nil, err
	}

	var list []*pb.ProductResponse

	for _, v := range products {
		u := pb.ProductResponse{
			Productid:  int32(v.ID),
			SellerId:   int32(v.SellerID),
			Name:       v.Name,
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

func (se *SellerService) GetProductByID(ctx context.Context, in *pb.GetProductByIDRequest) (*pb.ProductResponse, error) {
	product, err := se.SellerRepository.ReadProductID(int(in.ProductId))
	if err != nil {
		return nil, err
	}

	var response = &pb.ProductResponse{
		Productid:  int32(product.ID),
		SellerId:   int32(product.SellerID),
		Price:      float32(product.Price),
		Name:       product.Name,
		Stock:      int32(product.Stock),
		CategoryId: int32(product.Category_id),
		Discount:   int32(product.Discount),
	}

	return response, nil
}

func (se *SellerService) GetProductsByCategory(ctx context.Context, in *pb.GetProductByCategoryRequest) (*pb.GetProductsResponse, error) {
	products, err := se.SellerRepository.ReadProductCategory(in.CategoryName)
	if err != nil {
		return nil, err
	}

	var list []*pb.ProductResponse

	for _, v := range products {
		u := pb.ProductResponse{
			Productid:  int32(v.ID),
			SellerId:   int32(v.SellerID),
			Name:       v.Name,
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

func (se *SellerService) DeleteProduct(ctx context.Context, in *pb.DeleteProductRequest) (*emptypb.Empty, error) {
	err := se.SellerRepository.DeleteProduct(int(in.Productid), int(in.SellerId))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (se *SellerService) UpdateProduct(ctx context.Context, in *pb.UpdateProductRequest) (*pb.ProductResponse, error) {
	var product = model.Product{
		SellerID:    int(in.SellerId),
		Name:        in.Name,
		Price:       in.Price,
		Stock:       int(in.Stock),
		Category_id: int(in.CategoryId),
		Discount:    int(in.Discount),
	}

	err := se.SellerRepository.UpdateProduct(&product)
	if err != nil {
		return nil, err
	}

	var response = &pb.ProductResponse{
		Productid:  int32(product.ID),
		SellerId:   int32(product.SellerID),
		Name:       product.Name,
		Price:      float32(product.Price),
		Stock:      int32(product.Stock),
		CategoryId: int32(product.Category_id),
		Discount:   int32(product.Discount),
	}

	return response, nil
}

// seller service
func (se *SellerService) AddSeller(ctx context.Context, in *pb.AddSellerRequest) (*pb.SellerResponse, error) {
	var sellerInput = model.Seller{
		ID:         int(in.SellerId), // same as user id
		Name:       in.Name,
		LastActive: "",
	}

	seller, err := se.SellerRepository.CreateSeller(&sellerInput)
	if err != nil {
		return nil, err
	}

	var response = &pb.SellerResponse{
		SellerId:   int32(seller.ID),
		Name:       seller.Name,
		LastActive: seller.LastActive,
		AddressId:  0,
	}
	return response, nil
}

func (se *SellerService) AddAddress(ctx context.Context, in *pb.AddSellerAddressRequest) (*pb.AddressResponse, error) {
	var addressInput = model.Address{
		Name:    in.AddressName,
		Regency: in.AddressRegency,
		City:    in.AddressCity,
	}

	address, err := se.SellerRepository.CreateAddress(&addressInput)
	if err != nil {
		return nil, err
	}

	// update seller
	err = se.SellerRepository.UpdateAddressID(address.ID, int(in.SellerId))
	if err != nil {
		return nil, err
	}

	var response = &pb.AddressResponse{
		AddressId:      int32(address.ID),
		AddressName:    address.Name,
		AddressRegency: address.Regency,
		AddressCity:    address.City,
	}

	return response, nil
}

func (se *SellerService) GetAllSellers(ctx context.Context, in *emptypb.Empty) (*pb.GetSellersResponse, error) {
	sellers, err := se.SellerRepository.ReadAllSellers()
	if err != nil {
		return nil, err
	}

	var list []*pb.SellerResponse

	for _, v := range sellers {
		u := pb.SellerResponse{
			SellerId:   int32(v.ID),
			Name:       v.Name,
			LastActive: v.LastActive,
			AddressId:  int32(v.AddressID),
		}
		list = append(list, &u)
	}

	return &pb.GetSellersResponse{
		Sellers: list,
	}, nil
}

func (se *SellerService) GetSellerByID(ctx context.Context, in *pb.GetSellerByIDRequest) (*pb.SellerDetailResponse, error) {
	seller, err := se.SellerRepository.ReadSellerID(int(in.SellerId))
	if err != nil {
		return nil, err
	}

	var response = &pb.SellerDetailResponse{
		SellerId:       int32(seller.ID),
		Name:           seller.Name,
		LastActive:     seller.LastActive,
		AddressId:      int32(seller.Address.ID),
		AddressName:    seller.Address.Name,
		AddressRegency: seller.Address.Regency,
		AddressCity:    seller.Address.City,
	}

	return response, nil
}

func (se *SellerService) GetSellerByName(ctx context.Context, in *pb.GetSellerByNameRequest) (*pb.SellerDetailResponse, error) {
	seller, err := se.SellerRepository.ReadSellerName(in.Name)
	if err != nil {
		return nil, err
	}

	var response = &pb.SellerDetailResponse{
		SellerId:       int32(seller.ID),
		Name:           seller.Name,
		LastActive:     seller.LastActive,
		AddressId:      int32(seller.Address.ID),
		AddressName:    seller.Address.Name,
		AddressRegency: seller.Address.Regency,
		AddressCity:    seller.Address.City,
	}

	return response, nil
}

func (se *SellerService) UpdateAddress(ctx context.Context, in *pb.UpdateSellerAddressRequest) (*pb.AddressResponse, error) {
	var input = model.Address{
		ID:      int(in.AddressId),
		Name:    in.AddressName,
		Regency: in.AddressRegency,
		City:    in.AddressCity,
	}

	err := se.SellerRepository.UpdateAddress(&input)
	if err != nil {
		return nil, err
	}

	var response = &pb.AddressResponse{
		AddressId:      int32(input.ID),
		AddressName:    in.AddressName,
		AddressRegency: in.AddressRegency,
		AddressCity:    in.AddressCity,
	}

	return response, nil
}

func (se *SellerService) UpdateSellerName(ctx context.Context, in *pb.UpdateSellerNameRequest) (*emptypb.Empty, error) {
	err := se.SellerRepository.UpdateName(in.Name, int(in.SellerId))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (se *SellerService) UpdateSellerActivity(ctx context.Context, in *pb.UpdateSellerActivityRequest) (*emptypb.Empty, error) {
	now := time.Now().Format("2006-01-02 15:04:05")

	err := se.SellerRepository.UpdateActivity(now, int(in.SellerId))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
