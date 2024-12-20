package service

import (
	"context"

	home "gomall/app/frontend/hertz_gen/frontend/home"
	"gomall/app/frontend/infra/rpc"
	"gomall/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *home.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// resp := make(map[string]any)
	// items := []map[string]any{
	// 	{"Name": "T-shirt-1", "Price": 100, "Picture": "/static/image/t-shirt-1.jpeg"},
	// 	{"Name": "T-shirt-2", "Price": 110, "Picture": "/static/image/t-shirt-1.jpeg"},
	// 	{"Name": "T-shirt-3", "Price": 120, "Picture": "/static/image/t-shirt-2.jpeg"},
	// 	{"Name": "T-shirt-4", "Price": 130, "Picture": "/static/image/notebook.jpeg"},
	// 	{"Name": "T-shirt-5", "Price": 140, "Picture": "/static/image/t-shirt-1.jpeg"},
	// 	{"Name": "T-shirt-6", "Price": 150, "Picture": "/static/image/t-shirt.jpeg"},
	// }
	// resp["Title"] = "Hot Sales"
	// resp["Items"] = items
	// return resp, nil
	products, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"title": "Hot sale",
		"items": products.Products,
	}, nil
}
