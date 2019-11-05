package jaeger

import (
	"gin_api/util/jaeger_trace"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/envy"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func SetUp() gin.HandlerFunc {

	return func(c *gin.Context) {

		if envy.Get("JAEGER_ENABLE", "true") == "true" {

			var parentSpan opentracing.Span

			tracer, closer := jaeger_trace.NewJaegerTracer(envy.Get("APP_NAME", "app"), envy.Get("JAEGER_HOST_PORT","127.0.0.1:6831"))
			defer closer.Close()

			spCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
			if err != nil {
				parentSpan = tracer.StartSpan(c.Request.URL.Path)
				defer parentSpan.Finish()
			} else {
				parentSpan = opentracing.StartSpan(
					c.Request.URL.Path,
					opentracing.ChildOf(spCtx),
					opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
					ext.SpanKindRPCServer,
				)
				defer parentSpan.Finish()
			}
			c.Set("Tracer", tracer)
			c.Set("ParentSpanContext", parentSpan.Context())
		}
		c.Next()
	}
}
