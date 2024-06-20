package api

import (
	"net/http"

	_ "backend_course/branch_api_gateway/api/docs" //for swagger
	"backend_course/branch_api_gateway/api/handler"
	"backend_course/branch_api_gateway/config"
	"backend_course/branch_api_gateway/pkg/grpc_client"
	"backend_course/branch_api_gateway/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Config ...
type Config struct {
	Logger     logger.Logger
	GrpcClient *grpc_client.GrpcClient
	Cfg        config.Config
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(cnf Config) *gin.Engine {
	r := gin.New()

	r.Static("/images", "./static/images")

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "*")
	// config.AllowOrigins = cnf.Cfg.AllowOrigins
	r.Use(cors.New(config))

	handler := handler.New(&handler.HandlerConfig{
		Logger:     cnf.Logger,
		GrpcClient: cnf.GrpcClient,
		Cfg:        cnf.Cfg,
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Api gateway"})
	})

	r.POST("/v1/branch/create", handler.CreateBranch)
	r.GET("/v1/branch/getbyid/:id", handler.GetByID)
	r.GET("/v1/branch/getlist", handler.GetListBranch)
	r.PUT("/v1/branch/updatebranch", handler.Update)
	r.DELETE("/v1/branch/delete", handler.Delete)

	r.POST("/v1/customer/create", handler.CreateCustomer)
	r.GET("/v1/customer/getbyid/:id", handler.GetByIdCustomer)
	r.GET("/v1/customer/getlist", handler.GetListCustomer)
	r.PUT("/v1/customer/update", handler.UpdateCustomer)
	r.DELETE("/v1/customer/delete", handler.DeleteCustomer)

	r.POST("/v1/shop/create", handler.CreateShop)
	r.GET("/v1/shop/getbyid/:id", handler.GetByIdShop)
	r.GET("/v1/shop/getlist", handler.GetListShop)
	r.PUT("/v1/shop/update", handler.UpdateShop)
	r.DELETE("/v1/shop/delete", handler.DeleteShop)

	r.POST("/v1/seller/create", handler.CreateSeller)
	r.GET("/v1/seller/getbyid/:id", handler.GetByIdSeller)
	r.GET("/v1/seller/getlist", handler.GetListSeller)
	r.PUT("/v1/seller/update", handler.UpdateSeller)
	r.DELETE("/v1/seller/delete", handler.DeleteSeller)

	r.POST("/v1/system-user/create", handler.CreateSystemUser)
	r.GET("/v1/system-user/getbyid/:id", handler.GetByIdSystemUser)
	r.GET("/v1/system-user/getlist", handler.GetListSystemUser)
	r.PUT("/v1/system-user/update", handler.UpdateSystemUser)
	r.DELETE("/v1/system-user/delete", handler.DeleteSystemUser)

	// Shipper endpoints
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
