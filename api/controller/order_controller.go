package controller

import (
	"github.com/gin-gonic/gin"
	baseResponse "github.com/mariotj/olshop-app/api/entity/response"
	"github.com/mariotj/olshop-app/api/libraries/util"
	"github.com/mariotj/olshop-app/api/service"
	"net/http"
	"time"
)

func getOrderService() service.OrderService {
	return service.NewOrderService()
}

var orderService = getOrderService()

func GetDisplayOrderDetail(c *gin.Context)  {
	id := util.StrToUint(c.Params.ByName("id"))
	displayOrder, err := orderService.DisplayOrderDetails(id)

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = err.Error()

		c.JSON(http.StatusBadRequest, response)
	} else {
		response.Code = http.StatusOK
		response.Message = http.StatusText(http.StatusOK)
		response.Data = displayOrder

		c.JSON(http.StatusOK, response)
	}
}
