package service

import (
	"context"
	"fmt"
	"os"

	"gomall/app/checkout/infra/rpc"
	"gomall/rpc_gen/kitex_gen/cart"
	checkout "gomall/rpc_gen/kitex_gen/checkout"
	"gomall/rpc_gen/kitex_gen/email"
	"gomall/rpc_gen/kitex_gen/order"
	"gomall/rpc_gen/kitex_gen/payment"
	"gomall/rpc_gen/kitex_gen/product"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"google.golang.org/protobuf/proto"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}
	if cartResult == nil || cartResult.Items == nil {
		return nil, kerrors.NewGRPCBizStatusError(5004001, "cart is empty")
	}

	var total float32
	var oi []*order.OrderItem

	for _, cartItem := range cartResult.Items {
		productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{
			Id: cartItem.ProductId,
		})

		if resultErr != nil {
			return nil, resultErr
		}

		if productResp.Product == nil {
			continue
		}

		p := productResp.Product.Price

		cost := p * float32(cartItem.Quantity)
		total += cost

		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{
				ProductId: cartItem.ProductId,
				Quantity:  cartItem.Quantity,
			},
			Cost: cost,
		})
	}

	var orderId string

	orderResp, err := rpc.OrderClient.PlaceOrder(s.ctx, &order.PlaceOrderReq{
		Id:    req.UserId,
		Email: req.Email,
		Address: &checkout.Address{
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.Country,
			ZipCode:       req.Address.ZipCode,
		},
		Items: oi,
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5004002, err.Error())
	}

	if orderResp != nil && orderResp.Order != nil {
		orderId = orderResp.Order.OrderId
	}

	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderId,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
		},
	}

	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		klog.Error(err.Error())
	}

	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		return nil, err
	}
	// data, _ := proto.Marshal(&email.EmailReq{
	// 	From:        "from@example.com",
	// 	To:          req.Email,
	// 	ContentType: "text/plain",
	// 	Subject:     "You just created an order in CloudWeGo shop",
	// 	Content:     "You just created an order in CloudWeGo shop",
	// })

	// SendSyncMessage("string : You just created an order in CloudWeGo shop")
	data, _ := proto.Marshal(&email.EmailReq{
		From:        "from@example.com",
		To:          req.Email,
		ContentType: "text/plain",
		Subject:     "You just created an order in CloudWeGo shop",
		Content:     "You just created an order in CloudWeGo shop",
	})

	p, _ := rocketmq.NewProducer(
		// 设置  nameSrvAddr
		// nameSrvAddr 是 Topic 路由注册中心
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
		// 指定发送失败时的重试时间
		producer.WithRetry(2),
		// 设置 Group
		producer.WithGroupName("emailGroup"),
	)
	// 开始连接
	err = p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}

	// 设置节点名称
	topic := "email"
	// 循坏发送信息 (同步发送)

	msg := &primitive.Message{
		Topic: topic,
		Body:  data,
	}
	// 发送信息
	res, err := p.SendSync(context.Background(), msg)
	if err != nil {
		fmt.Printf("send message error:%s\n", err)
	} else {
		fmt.Printf("send message success: result=%s\n", res.String())
	}

	// 关闭生产者
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error:%s", err.Error())
	}

	klog.Info(paymentResult)

	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId,
	}
	return
}

// 生产者发送数据
// func SendSyncMessage(message string) {
// 	endPoint := []string{"127.0.0.1:9876"}
// 	// 创建一个producer实例
// 	p, _ := rocketmq.NewProducer(
// 		producer.WithNameServer(endPoint),            // 服务器地址
// 		producer.WithRetry(2),                        // 尝试发送数据的次数
// 		producer.WithGroupName("EmailProducerGroup"), // 生产者组的名称
// 	)
// 	// 启动生产者
// 	err := p.Start()
// 	// 如果启动失败，则退出
// 	if err != nil {
// 		fmt.Printf("start producer error: %s", err.Error())
// 		os.Exit(1)
// 	}

// 	// 发送消息，sync同步发送，async异步发送
// 	result, err := p.SendSync(context.Background(),
// 		&primitive.Message{
// 			Topic: "email",
// 			Body:  []byte(message),
// 		})
// 	// 检测msg师傅哦
// 	if err != nil {
// 		fmt.Printf("send message error: %s\n", err.Error())
// 	} else {
// 		fmt.Printf("send message seccess: result=%s\n", result.String())
// 	}
// }

// func PreduceInit() {
// 	data, _ := proto.Marshal(&email.EmailReq{
// 		From:        "from@example.com",
// 		To:          req.Email,
// 		ContentType: "text/plain",
// 		Subject:     "You just created an order in CloudWeGo shop",
// 		Content:     "You just created an order in CloudWeGo shop",
// 	})

// 	p, _ := rocketmq.NewProducer(
// 		// 设置  nameSrvAddr
// 		// nameSrvAddr 是 Topic 路由注册中心
// 		producer.WithNameServer([]string{"127.0.0.1:9876"}),
// 		// 指定发送失败时的重试时间
// 		producer.WithRetry(2),
// 		// 设置 Group
// 		producer.WithGroupName("emailGroup"),
// 	)
// 	// 开始连接
// 	err := p.Start()
// 	if err != nil {
// 		fmt.Printf("start producer error: %s", err.Error())
// 		os.Exit(1)
// 	}

// 	// 设置节点名称
// 	topic := "email"
// 	// 循坏发送信息 (同步发送)

// 	msg := &primitive.Message{
// 		Topic: topic,
// 		Body:  data,
// 	}
// 	// 发送信息
// 	res, err := p.SendSync(context.Background(), msg)
// 	if err != nil {
// 		fmt.Printf("send message error:%s\n", err)
// 	} else {
// 		fmt.Printf("send message success: result=%s\n", res.String())
// 	}

// 	// 关闭生产者
// 	err = p.Shutdown()
// 	if err != nil {
// 		fmt.Printf("shutdown producer error:%s", err.Error())
// 	}
// }
