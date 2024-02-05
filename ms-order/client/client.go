package client

import (
	"ms-order/config"
	sellerpb "ms-seller/pb"
	userpb "ms-user/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	Payment PaymentClient
	Seller  MsSellerClient
	User    MsUserClient
}

type NewClientParam struct {
	Cfg        config.Config
	SellerConn *grpc.ClientConn
	UserConn   *grpc.ClientConn
}

func New(cfg config.Config) (*Client, func(), error) {
	sellerconn, err := grpc.Dial(":50002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	userconn, err := grpc.Dial(":50001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	closefn := func() {
		sellerconn.Close()
		userconn.Close()
	}

	client := &Client{
		Payment: PaymentClient{cfg: cfg},
		Seller:  MsSellerClient{client: sellerpb.NewSellerServiceClient(sellerconn)},
		User:    MsUserClient{client: userpb.NewUserServiceClient(userconn)},
	}
	return client, closefn, nil
}
