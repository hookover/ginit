package grpc_client

import (
	"gin_api/util/grpc_log"
	"gin_api/util/jaeger_trace"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/envy"
	grpc_middeware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"time"
)

func CreateServiceListenConn(c *gin.Context) *grpc.ClientConn {
	return createGrpcConn("127.0.0.1:9901", c)
}

func CreateServiceSpeakConn(c *gin.Context) *grpc.ClientConn {
	return createGrpcConn("127.0.0.1:9902", c)
}

func CreateServiceReadConn(c *gin.Context) *grpc.ClientConn {
	return createGrpcConn("127.0.0.1:9903", c)
}

func CreateServiceWriteConn(c *gin.Context) *grpc.ClientConn {
	return createGrpcConn("127.0.0.1:9904", c)
}

func createGrpcConn(serviceAddress string, c *gin.Context) *grpc.ClientConn {

	var conn *grpc.ClientConn
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()

	if envy.Get("JAEGER_ENABLE", "true") == "true" {

		tracer, _ := c.Get("Tracer")
		parentSpanContext, _ := c.Get("ParentSpanContext")

		conn, err = grpc.DialContext(
			ctx,
			serviceAddress,
			grpc.WithInsecure(),
			grpc.WithBlock(),
			grpc.WithUnaryInterceptor(
				grpc_middeware.ChainUnaryClient(
					jaeger_trace.ClientInterceptor(tracer.(opentracing.Tracer), parentSpanContext.(opentracing.SpanContext)),
					grpc_log.ClientInterceptor(),
				),
			),
		)
	} else {
		conn, err = grpc.DialContext(
			ctx,
			serviceAddress,
			grpc.WithInsecure(),
			grpc.WithBlock(),
			grpc.WithUnaryInterceptor(
				grpc_middeware.ChainUnaryClient(
					grpc_log.ClientInterceptor(),
				),
			),
		)
	}

	if err != nil {
		fmt.Println(serviceAddress, "grpc conn err:", err)
	}
	return conn
}
