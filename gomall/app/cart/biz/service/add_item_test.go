package service

import (
	"context"
	"testing"

	"gomall/app/cart/biz/dal/mysql"
	cart "gomall/rpc_gen/kitex_gen/cart"

	"github.com/joho/godotenv"
)

func TestAddItem_Run(t *testing.T) {
	godotenv.Load("D:\\golang_projects\\gomall\\app\\cart\\.env")
	mysql.Init()
	ctx := context.Background()
	s := NewAddItemService(ctx)
	// init req and assert value

	req := &cart.AddItemReq{
		UserId: 1,
		Item:   &cart.CartItem{ProductId: 1, Quantity: 1},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
