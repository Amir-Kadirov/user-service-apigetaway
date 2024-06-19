package grpc_client

import (
	pc "backend_course/branch_api_gateway/genproto/user_service"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"backend_course/branch_api_gateway/config"
)

// GrpcClientI ...
type GrpcClientI interface {
	BranchBranch() pc.BranchServiceClient
	CustomerService() pc.CustomerServiceClient
}

// GrpcClient ...
type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

// New ...
func New(cfg config.Config) (*GrpcClient, error) {

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.UserBranchHost, cfg.UserBranchPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("catalog service dial host: %s port:%s err: %s",
			cfg.UserBranchHost, cfg.UserBranchPort, err)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"branch_service":   pc.NewBranchServiceClient(connUser),
			"customer_service": pc.NewCustomerServiceClient(connUser),
			"shop_service":     pc.NewShopServiceClient(connUser),
			"seller_service":   pc.NewSellerServiceClient(connUser),
			"system_user_service": pc.NewSystemUserServiceClient(connUser),
		},
	}, nil
}

func (g *GrpcClient) BranchBranch() pc.BranchServiceClient {
	return g.connections["branch_service"].(pc.BranchServiceClient)
}

func (g *GrpcClient) CustomerService() pc.CustomerServiceClient {
	return g.connections["customer_service"].(pc.CustomerServiceClient)
}

func (g *GrpcClient) ShopService() pc.ShopServiceClient {
	return g.connections["shop_service"].(pc.ShopServiceClient)
}

func (g *GrpcClient) SellerService() pc.SellerServiceClient {
	return g.connections["seller_service"].(pc.SellerServiceClient)
}

func (g *GrpcClient) SystemUserService() pc.SystemUserServiceClient {
	return g.connections["system_user_service"].(pc.SystemUserServiceClient)
}
