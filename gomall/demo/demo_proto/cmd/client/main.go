package main

import (
	"context"
	"fmt"
	"log"

	"gomall/demo/demo_proto/kitex_gen/pbapi"
	"gomall/demo/demo_proto/kitex_gen/pbapi/echo"
	"gomall/demo/demo_proto/middleware"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Fatal(err)
	}
	c, err := echo.NewClient("demo_proto", client.WithResolver(r),
		client.WithTransportProtocol(transport.GRPC), // win11不使用 gRPC
		// client.WithShortConnection(), // 使用短链接
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithMiddleware(middleware.Middleware))
	if err != nil {
		log.Fatal(err)
	}
	ctx := metainfo.WithPersistentValue(context.Background(), "CLIENT_NAME", "demo_proto_client")

	res, err := c.Echo(ctx, &pbapi.Request{Message: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", res)
}
