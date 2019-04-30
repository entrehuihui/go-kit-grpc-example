package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"../../ordering"
	"../operating"
	"../proto"
	"./pkg/ui/data/swagger"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	"github.com/maxiiot/maxiiot/logger"
	"google.golang.org/grpc"
)

type config struct {
	logLevel   string
	dsn        string
	httpPort   string
	grpcPort   string
	secret     string
	serverCert string
	serverKey  string
}

func loadConfig() config {
	return config{
		logLevel: "info",
		dsn:      "postgres://maxiiot:maxiiot@localhost:5432/maxiiot?sslmode=disable",
		httpPort: "8880",
		grpcPort: "8881",
		secret:   "verysecret",
	}
}

func main() {
	cfg := loadConfig()
	logger, err := logger.New(os.Stdout, cfg.logLevel)
	if err != nil {
		log.Fatalf(err.Error())
	}
	db, err := ordering.Connect(cfg.dsn)
	if err != nil {
		logger.Error(err.Error())
	}
	defer db.Close()
	errs := make(chan error, 2)

	go startGRPC(db, cfg, errs)
	go startGW(cfg, errs)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()
	err = <-errs
	logger.Error(fmt.Sprintf("ordering service terminated: %s", err))
}

func startGRPC(db *sqlx.DB, cfg config, errs chan error) {
	svc := operating.NewOrderoperating(db)
	service := ordering.NewInstrumentingService(svc)
	ls, _ := net.Listen("tcp", ":"+cfg.grpcPort)
	gs := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))

	proto.RegisterCommodityServer(gs, service)
	proto.RegisterOrderinfoServer(gs, service)
	proto.RegisterPayServer(gs, service)
	proto.RegisterRefundServer(gs, service)
	fmt.Println("GRPC Listen on : ", cfg.grpcPort)
	errs <- gs.Serve(ls)
}

func startGW(cfg config, errs chan error) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	// 商品
	err := proto.RegisterCommodityHandlerFromEndpoint(ctx, mux, ":"+cfg.grpcPort, opts)
	if err != nil {
		errs <- err
	}
	// 订单
	err = proto.RegisterOrderinfoHandlerFromEndpoint(ctx, mux, ":"+cfg.grpcPort, opts)
	if err != nil {
		errs <- err
	}
	// 付款
	err = proto.RegisterPayHandlerFromEndpoint(ctx, mux, ":"+cfg.grpcPort, opts)
	if err != nil {
		errs <- err
	}
	// 订单
	err = proto.RegisterRefundHandlerFromEndpoint(ctx, mux, ":"+cfg.grpcPort, opts)
	if err != nil {
		errs <- err
	}

	gwmux := http.NewServeMux()
	gwmux.Handle("/", mux)
	serveSwagger(gwmux)
	fmt.Println("GRPCGW Listen on : ", cfg.httpPort)
	errs <- http.ListenAndServe(":"+cfg.httpPort, gwmux)
}

//swagger UI
func serveSwagger(mux *http.ServeMux) {
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
		io.Copy(w, strings.NewReader(proto.AiasJSON))
	})

	mime.AddExtensionType(".svg", "image/svg+xml")

	// Expose files in third_party/swagger-ui/ on <host>/swagger-ui
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}
