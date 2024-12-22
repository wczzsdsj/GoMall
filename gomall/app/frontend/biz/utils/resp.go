package utils

import (
	"context"

	"gomall/app/frontend/infra/rpc"
	frontendUtils "gomall/app/frontend/utils"
	"gomall/rpc_gen/kitex_gen/cart"

	"github.com/cloudwego/hertz/pkg/app"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	// session := sessions.Default(c)
	// userId := session.Get("user_id")
	// content["user_id"] = userId

	// content["user_id"] = ctx.Value(middleware.SessionUserId)
	userId := frontendUtils.GetUserIdFromCtx(ctx)
	content["user_id"] = userId

	if userId > 0 {
		cartResp, err := rpc.CartClient.GetCart(ctx, &cart.GetCartReq{
			UserId: uint32(userId),
		})
		if err == nil && cartResp != nil {
			content["cart_num"] = len(cartResp.Items)
		}
	}
	return content
}
