package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mariotj/olshop-app/api/controller"
	"github.com/mariotj/olshop-app/api/entity/path"
	"github.com/mariotj/olshop-app/api/entity/response"
	"github.com/mariotj/olshop-app/api/service"
	"net/http"
	"time"
)

func Run() *gin.Engine {
	engine := gin.Default()
	engine.RedirectTrailingSlash = false

	service.Initialize()

	v1 := engine.Group(path.BaseUrl)
	{
		v1.GET(path.OrderDetailsById, controller.GetDisplayOrderDetail)
	}

	engine.NoRoute(func(context *gin.Context) {
		var resp = &response.BaseResponse{
			ServerTime:	time.Now(),
		}

		resp.Code = http.StatusNotFound
		resp.Message = "Route not found"

		context.JSON(http.StatusNotFound, resp)
	})

	return engine
}