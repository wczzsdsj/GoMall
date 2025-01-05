package mq

import (
	"context"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func CreateTopic(topicName string) {
	endPoint := []string{"127.0.0.1:9876"}
	// 创建主题
	// 先连接远程的服务器，得到一个具柄testAdmin，然后利用该具柄创建CreateTopic()创建topic
	testAdmin, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(endPoint)))
	// 检查是否连接成功
	if err != nil {
		fmt.Printf("connection error: %s\n", err.Error())
	}
	err = testAdmin.CreateTopic(context.Background(),
		admin.WithTopicCreate(topicName))
	// 检查是否创建topic失败
	if err != nil {
		fmt.Printf("createTopic error: %s\n", err.Error())
	}
}
