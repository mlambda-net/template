package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mlambda-net/template/pkg/infrastructure/endpoint/api"
	"github.com/mlambda-net/template/pkg/infrastructure/ports/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	services := api.NewApi()

	docs.SwaggerInfo.Title = "Template API"
	docs.SwaggerInfo.Description = "This is the api for the template service."
	docs.SwaggerInfo.Version = services.GetVersion()
	docs.SwaggerInfo.Host = services.GetHost()
	docs.SwaggerInfo.BasePath = services.Path()
	docs.SwaggerInfo.Schemes = []string{"http", "https"}


	go func() {
		r := gin.New()
		r.GET("/template/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    _ = r.Run( fmt.Sprintf(":%d", services.Docs()))
	}()

	services.Start()

}
