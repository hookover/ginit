module gin_api

go 1.12

require (
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/gin-contrib/pprof v1.2.1
	github.com/gin-gonic/gin v1.4.0
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.16.0 // indirect
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/xorm v0.7.9
	github.com/gobuffalo/envy v1.7.1
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0
	github.com/hookover/logging v0.0.3
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/opentracing/opentracing-go v1.1.0
	github.com/rs/zerolog v1.16.0
	github.com/uber-go/atomic v1.5.0 // indirect
	github.com/uber/jaeger-client-go v2.19.0+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	github.com/ugorji/go/codec v0.0.0-20181022190402-e5e69e061d4f // indirect
	golang.org/x/time v0.0.0-20181108054448-85acf8d2951c
	google.golang.org/grpc v1.19.0
	gopkg.in/go-playground/validator.v9 v9.30.0
	xorm.io/core v0.7.2-0.20190928055935-90aeac8d08eb
)

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43

replace github.com/uber-go/atomic v1.5.0 => go.uber.org/atomic v1.5.0
