package order

import (
	"time"

	validator "github.com/code-kakitai/go-pkg/validator"
	"github.com/gin-gonic/gin"

	orderApp "github.com/yumekumo/sauna-shop/application/order"
	"github.com/yumekumo/sauna-shop/presentation/settings"
)

type handler struct {
	saveOrderUseCase *orderApp.SaveOrderUseCase
}

func NewHandler(saveOrderUseCase *orderApp.SaveOrderUseCase) handler {
	return handler{
		saveOrderUseCase: saveOrderUseCase,
	}
}

// PostOrders godoc
// @Summary 注文をする
// @Tags orders
// @Accept json
// @Produce json
// @Param request body []PostOrdersParams true "注文商品"
// @Success 200 {int} id
// @Router /v1/orders [post]
func (h handler) PostOrders(ctx *gin.Context) {
	var params []*PostOrdersParams
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		settings.ReturnBadRequest(ctx, err)
		return
	}
	validate := validator.GetValidator()
	if err := validate.Struct(&params); err != nil {
		settings.ReturnStatusBadRequest(ctx, err)
		return
	}

	// 本来はsessionに入っているuserIDを取得するが、本質ではないため省略
	userID := "01HCNYK0PKYZWB0ZT1KR0EPWGP"
	dtos := make([]orderApp.SaveOrderUseCaseInputDto, 0, len(params))

	for _, param := range params {
		dtos = append(dtos, orderApp.SaveOrderUseCaseInputDto{
			ProductID: param.ProductID,
			Quantity:  param.Quantity,
		})
	}
	id, err := h.saveOrderUseCase.Run(
		ctx.Request.Context(),
		userID,
		dtos,
		time.Now(),
	)
	if err != nil {
		settings.ReturnStatusInternalServerError(ctx, err)
		return
	}

	settings.ReturnStatusCreated(ctx, id)
}
