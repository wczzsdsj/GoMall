package service

import (
	"context"

	"gomall/app/product/biz/dal/mysql"
	"gomall/app/product/biz/model"
	product "gomall/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	if req.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "product id is required")
	}
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)

	p, err := productQuery.GetById(int(req.Id))
	if err != nil {
		return nil, err
	}

	return &product.GetProductResp{
		Product: &product.Product{
			Id:          uint32(p.ID),
			Picture:     p.Picture,
			Price:       p.Price,
			Description: p.Description,
			Name:        p.Name,
		},
	}, nil
}
