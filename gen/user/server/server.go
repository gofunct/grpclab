package server

import (
	"fmt"
	"github.com/go-kit/kit/log"
	pb "github.com/gofunct/grpclab/gen/user"
	"github.com/gofunct/grpclab/gen/user/endpoints"
	grpctransport "github.com/gofunct/grpclab/gen/user/transports/grpc"
	httptransport "github.com/gofunct/grpclab/gen/user/transports/http"
	svc "github.com/gofunct/grpclab/services/user"
	"github.com/gorilla/handlers"
	"github.com/spf13/viper"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"google.golang.org/grpc"
)

func RunServer(v *viper.Viper, logger log.Logger) {
	mux := http.NewServeMux()
	errc := make(chan error)
	s := grpc.NewServer()

	// initialize services here...
	{
		svc := svc.New()
		endpoints := endpoints.MakeEndpoints(svc)
		srv := grpctransport.MakeGRPCServer(endpoints)
		pb.RegisterUserServiceServer(s, srv)
		httptransport.RegisterHandlers(svc, mux, endpoints)
	}

	// start servers
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger := log.With(logger, "transport", "HTTP")
		logger.Log("addr", v.GetString("http_port"))
		errc <- http.ListenAndServe(v.GetString("http_port"), handlers.LoggingHandler(os.Stderr, mux))
	}()

	go func() {
		logger := log.With(logger, "transport", "gRPC")
		ln, err := net.Listen("tcp", v.GetString("grpc_port"))
		if err != nil {
			errc <- err
			return
		}
		logger.Log("addr", v.GetString("grpc_port"))
		errc <- s.Serve(ln)
	}()

	logger.Log("exit", <-errc)
}

