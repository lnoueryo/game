package config

import (
	"log"
	"os"
	"time"
	"google.golang.org/grpc"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Infolog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
var Errorlog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

// func CreateAccesslog() {
// 	start := time.Now()
// 	Infolog.Printf("%s %s %v %v", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
// }
func CreateAccesslog() grpc.ServerOption {
	zapLogger, _ := zap.NewProduction()
	opts := []grpc_zap.Option{
		grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
				return zap.Int64("grpc.time_ns", duration.Nanoseconds())
		}),
	}
	grpc_zap.ReplaceGrpcLogger(zapLogger)
	return grpc_middleware.WithUnaryServerChain(
		grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		grpc_zap.UnaryServerInterceptor(zapLogger, opts...),
	)
}