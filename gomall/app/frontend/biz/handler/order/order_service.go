package order

import (
	"context"

	"gomall/app/frontend/biz/service"
	"gomall/app/frontend/biz/utils"
	common "gomall/app/frontend/hertz_gen/frontend/common"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// OrderList .
// @router /order [GET]
func OrderList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// resp := &common.Empty{}
	resp, err := service.NewOrderListService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
	c.HTML(consts.StatusOK, "order", utils.WarpResponse(ctx, c, resp))
}
