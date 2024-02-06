package client

import sellerpb "ms-seller/pb"

type MsSellerClient struct {
	client sellerpb.SellerServiceClient
}
