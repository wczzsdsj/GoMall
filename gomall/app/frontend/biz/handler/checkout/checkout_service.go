package checkout

import (
	"context"

	"gomall/app/frontend/biz/service"
	"gomall/app/frontend/biz/utils"
	checkout "gomall/app/frontend/hertz_gen/frontend/checkout"
	common "gomall/app/frontend/hertz_gen/frontend/common"

	hertzUtils "github.com/cloudwego/hertz/pkg/common/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Checkout .
// @router /checkout [GET]
func Checkout(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// resp := &common.Empty{}
	resp, err := service.NewCheckoutService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "checkout", utils.WarpResponse(ctx, c, resp))
}

// CheckoutWaiting .
// @router /checkout/waiting [POST]
func CheckoutWaiting(ctx context.Context, c *app.RequestContext) {
	var err error
	var req checkout.CheckoutReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.HTML(consts.StatusOK, "waiting", utils.WarpResponse(ctx, c, hertzUtils.H{"warning": err}))
	}

	// resp := &common.Empty{}
	resp, err := service.NewCheckoutWaitingService(ctx, c).Run(&req)
	if err != nil {
		c.HTML(consts.StatusOK, "waiting", utils.WarpResponse(ctx, c, hertzUtils.H{"error": err}))
		return
	}

	c.HTML(consts.StatusOK, "waiting", utils.WarpResponse(ctx, c, resp))
}

// CheckoutResult .
// @router /checkout/result [GET]
func CheckoutResult(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		c.HTML(consts.StatusOK, "waiting", utils.WarpResponse(ctx, c, hertzUtils.H{"warning": err}))
		return
	}

	// resp := &common.Empty{}
	resp, err := service.NewCheckoutResultService(ctx, c).Run(&req)
	if err != nil {
		c.HTML(consts.StatusOK, "waiting", utils.WarpResponse(ctx, c, hertzUtils.H{"error": err}))
		return
	}

	c.HTML(consts.StatusOK, "result", utils.WarpResponse(ctx, c, resp))
}
