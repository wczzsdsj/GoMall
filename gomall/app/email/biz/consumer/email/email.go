package email

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"gomall/rpc_gen/kitex_gen/email"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"google.golang.org/protobuf/proto"
)

func ConsumerInit() {
	// // 创建消费者
	// c, err := consumer.NewPushConsumer(
	// 	consumer.WithGroupName("email_consumer_group"),
	// 	consumer.WithNameServer([]string{"127.0.0.1:9876"}), // 替换为您的 Name Server 地址
	// )
	// if err != nil {
	// 	log.Panic(err)
	// }
	// err = c.Subscribe("email", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	// 	for _, msg := range msgs {
	// 		log.Printf("Message ID: %s, Body: %s\n", msg.MsgId, string(msg.Body))
	// 		// 处理消息逻辑
	// 		var req email.EmailReq
	// 		err := proto.Unmarshal(msg.Body, &req)
	// 		if err != nil {
	// 			klog.Error(err)
	// 		}
	// 		noopEmail := notify.NewNoopEmail()
	// 		_ = noopEmail.Send(&req)
	// 	}
	// 	return consumer.ConsumeSuccess, nil
	// })
	// if err != nil {
	// 	fmt.Printf("subscribe message error: %s\n", err.Error())
	// }

	// // 启动消费者
	// err = c.Start()
	// if err != nil {
	// 	log.Panic(err)
	// }
	// log.Println("Consumer started")

	// // 阻塞主线程，保持运行
	// // select {}
	c, _ := rocketmq.NewPushConsumer(
		// 消费组
		consumer.WithGroupName("emailGroup"),
		// namesrv地址
		consumer.WithNameServer([]string{"127.0.0.1:9876"}),
	)
	// 必须先在 开始前
	err := c.Subscribe("email", consumer.MessageSelector{}, func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range ext {
			_, err := handleMessage(msg)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	err = c.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	time.Sleep(time.Hour)
	err = c.Shutdown()
	if err != nil {
		fmt.Printf("shutdown Consumer error:%s", err.Error())
	}
}

func handleMessage(msg *primitive.MessageExt) (consumer.ConsumeResult, error) {
	// 反序列化消息
	var emailReq email.EmailReq
	err := proto.Unmarshal(msg.Body, &emailReq)
	if err != nil {
		log.Printf("Failed to unmarshal email request: %v", err)
		return consumer.ConsumeRetryLater, nil // 可选择重试
	}

	// 处理反序列化后的消息
	log.Printf("Received email request: From=%s, To=%s, Subject=%s", emailReq.From, emailReq.To, emailReq.Subject)

	return consumer.ConsumeSuccess, nil
}
