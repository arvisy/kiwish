package client

import userpb "ms-user/pb"

type MsUserClient struct {
	client userpb.UserServiceClient
}
