package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/flowchartsman/swaggerui"
	"github.com/fullstorydev/grpcui/standalone"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/soheilhy/cmux"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	protobuf "google.golang.org/protobuf/proto"

	"northwind/proto"
	pb_Category "northwind/proto/category"
	pb_Customer "northwind/proto/customer"
	pb_CustomerCustomerDemo "northwind/proto/customer_customer_demo"
	pb_CustomerDemographic "northwind/proto/customer_demographic"
	pb_Employee "northwind/proto/employee"
	pb_EmployeeTerritory "northwind/proto/employee_territory"
	pb_Order "northwind/proto/order"
	pb_OrderDetail "northwind/proto/order_detail"
	pb_Product "northwind/proto/product"
	pb_Region "northwind/proto/region"
	pb_Shipper "northwind/proto/shipper"
	pb_Supplier "northwind/proto/supplier"
	pb_Territory "northwind/proto/territory"
	pb_UsState "northwind/proto/us_state"
)

const readTimeout = 15 * time.Second

// Server represents a gRPC server
type Server struct {
	cfg                         Config
	log                         *zap.Logger
	grpcServer                  *grpc.Server
	healthServer                *health.Server
	CategoryService             pb_Category.CategoryServer
	CustomerService             pb_Customer.CustomerServer
	CustomerCustomerDemoService pb_CustomerCustomerDemo.CustomerCustomerDemoServer
	CustomerDemographicService  pb_CustomerDemographic.CustomerDemographicServer
	EmployeeService             pb_Employee.EmployeeServer
	EmployeeTerritoryService    pb_EmployeeTerritory.EmployeeTerritoryServer
	OrderService                pb_Order.OrderServer
	OrderDetailService          pb_OrderDetail.OrderDetailServer
	ProductService              pb_Product.ProductServer
	RegionService               pb_Region.RegionServer
	ShipperService              pb_Shipper.ShipperServer
	SupplierService             pb_Supplier.SupplierServer
	TerritoryService            pb_Territory.TerritoryServer
	UsStateService              pb_UsState.UsStateServer
}

// New gRPC server
func New(cfg Config, log *zap.Logger,
	CategoryService pb_Category.CategoryServer,
	CustomerService pb_Customer.CustomerServer,
	CustomerCustomerDemoService pb_CustomerCustomerDemo.CustomerCustomerDemoServer,
	CustomerDemographicService pb_CustomerDemographic.CustomerDemographicServer,
	EmployeeService pb_Employee.EmployeeServer,
	EmployeeTerritoryService pb_EmployeeTerritory.EmployeeTerritoryServer,
	OrderService pb_Order.OrderServer,
	OrderDetailService pb_OrderDetail.OrderDetailServer,
	ProductService pb_Product.ProductServer,
	RegionService pb_Region.RegionServer,
	ShipperService pb_Shipper.ShipperServer,
	SupplierService pb_Supplier.SupplierServer,
	TerritoryService pb_Territory.TerritoryServer,
	UsStateService pb_UsState.UsStateServer,

) *Server {
	return &Server{
		cfg:                         cfg,
		log:                         log,
		CategoryService:             CategoryService,
		CustomerService:             CustomerService,
		CustomerCustomerDemoService: CustomerCustomerDemoService,
		CustomerDemographicService:  CustomerDemographicService,
		EmployeeService:             EmployeeService,
		EmployeeTerritoryService:    EmployeeTerritoryService,
		OrderService:                OrderService,
		OrderDetailService:          OrderDetailService,
		ProductService:              ProductService,
		RegionService:               RegionService,
		ShipperService:              ShipperService,
		SupplierService:             SupplierService,
		TerritoryService:            TerritoryService,
		UsStateService:              UsStateService,
	}
}

// ListenAndServe start the server
func (srv *Server) ListenAndServe() error {
	grpc_zap.ReplaceGrpcLoggerV2(srv.log)
	srv.grpcServer = grpc.NewServer(srv.cfg.grpcOpts(srv.log)...)
	reflection.Register(srv.grpcServer)
	srv.healthServer = health.NewServer()
	healthpb.RegisterHealthServer(srv.grpcServer, srv.healthServer)
	srv.healthServer.SetServingStatus("ww", healthpb.HealthCheckResponse_SERVING)

	pb_Category.RegisterCategoryServer(srv.grpcServer, srv.CategoryService)
	pb_Customer.RegisterCustomerServer(srv.grpcServer, srv.CustomerService)
	pb_CustomerCustomerDemo.RegisterCustomerCustomerDemoServer(srv.grpcServer, srv.CustomerCustomerDemoService)
	pb_CustomerDemographic.RegisterCustomerDemographicServer(srv.grpcServer, srv.CustomerDemographicService)
	pb_Employee.RegisterEmployeeServer(srv.grpcServer, srv.EmployeeService)
	pb_EmployeeTerritory.RegisterEmployeeTerritoryServer(srv.grpcServer, srv.EmployeeTerritoryService)
	pb_Order.RegisterOrderServer(srv.grpcServer, srv.OrderService)
	pb_OrderDetail.RegisterOrderDetailServer(srv.grpcServer, srv.OrderDetailService)
	pb_Product.RegisterProductServer(srv.grpcServer, srv.ProductService)
	pb_Region.RegisterRegionServer(srv.grpcServer, srv.RegionService)
	pb_Shipper.RegisterShipperServer(srv.grpcServer, srv.ShipperService)
	pb_Supplier.RegisterSupplierServer(srv.grpcServer, srv.SupplierService)
	pb_Territory.RegisterTerritoryServer(srv.grpcServer, srv.TerritoryService)
	pb_UsState.RegisterUsStateServer(srv.grpcServer, srv.UsStateService)

	var listen net.Listener
	dialOptions := []grpc.DialOption{grpc.WithBlock()}
	var schema string
	if srv.cfg.TLSEnabled() {
		schema = "https"
		tlsCert, err := tls.LoadX509KeyPair(srv.cfg.Cert, srv.cfg.Key)
		if err != nil {
			return fmt.Errorf("failed to parse certificate and key: %w", err)
		}
		tlsCert.Leaf, _ = x509.ParseCertificate(tlsCert.Certificate[0])
		tc := &tls.Config{
			Certificates: []tls.Certificate{tlsCert},
			MinVersion:   tls.VersionTLS12,
		}
		listen, err = tls.Listen("tcp", fmt.Sprintf(":%d", srv.cfg.Port), tc)
		if err != nil {
			return err
		}

		cp := x509.NewCertPool()
		cp.AddCert(tlsCert.Leaf)
		dialOptions = append(dialOptions, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(cp, "")))
	} else {
		schema = "http"
		var err error
		listen, err = net.Listen("tcp", fmt.Sprintf(":%d", srv.cfg.Port))
		if err != nil {
			return err
		}
		dialOptions = append(dialOptions, grpc.WithInsecure())
	}

	mux := cmux.New(listen)
	mux.SetReadTimeout(readTimeout)
	grpcListener := mux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpListener := mux.Match(cmux.Any())

	go func() {
		if err := mux.Serve(); err != nil {
			srv.log.Error("Failed to serve cmux", zap.String("error", err.Error()))
		}
	}()

	if srv.cfg.PrometheusEnabled() {
		grpc_prometheus.Register(srv.grpcServer)
		go prometheusServer(srv.log, srv.cfg.PrometheusPort)
	}

	go func() {
		srv.log.Info("Server running", zap.String("addr", grpcListener.Addr().String()))
		if err := srv.grpcServer.Serve(grpcListener); err != nil {
			srv.log.Fatal("Failed to start gRPC Server", zap.String("error", err.Error()))
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()

	sAddr := fmt.Sprintf("dns:///localhost:%d", srv.cfg.Port)
	cc, err := grpc.DialContext(
		ctx,
		sAddr,
		dialOptions...,
	)
	if err != nil {
		return err
	}
	defer cc.Close()

	gwmux := runtime.NewServeMux(
		runtime.WithMetadata(annotator),
		runtime.WithForwardResponseOption(forwardResponse),
		runtime.WithOutgoingHeaderMatcher(outcomingHeaderMatcher),
	)
	if err := pb_Category.RegisterCategoryHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}
	if err := pb_Customer.RegisterCustomerHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}
	if err := pb_CustomerCustomerDemo.RegisterCustomerCustomerDemoHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}
	if err := pb_CustomerDemographic.RegisterCustomerDemographicHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}
	if err := pb_Employee.RegisterEmployeeHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}
	if err := pb_EmployeeTerritory.RegisterEmployeeTerritoryHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}
	if err := pb_Order.RegisterOrderHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}
	if err := pb_OrderDetail.RegisterOrderDetailHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}
	if err := pb_Product.RegisterProductHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}
	if err := pb_Region.RegisterRegionHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}
	if err := pb_Shipper.RegisterShipperHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}
	if err := pb_Supplier.RegisterSupplierHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}
	if err := pb_Territory.RegisterTerritoryHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}
	if err := pb_UsState.RegisterUsStateHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}

	httpMux := http.NewServeMux()

	grpcui, err := standalone.HandlerViaReflection(ctx, cc, sAddr)
	if err != nil {
		return err
	}

	httpMux.Handle("/grpcui/", http.StripPrefix("/grpcui", grpcui))
	httpMux.Handle("/swagger/", http.StripPrefix("/swagger", swaggerui.Handler(proto.OpenAPIv2)))
	httpMux.Handle("/", gwmux)

	httpServer := &http.Server{
		Handler: httpMux,
	}

	srv.log.Info(fmt.Sprintf("Serving gRPC UI on %s://localhost:%d/grpcui", schema, srv.cfg.Port))
	srv.log.Info(fmt.Sprintf("Serving Swagger UI on %s://localhost:%d/swagger", schema, srv.cfg.Port))
	return httpServer.Serve(httpListener)
}

func prometheusServer(log *zap.Logger, port int) {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  readTimeout,
		WriteTimeout: time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      mux,
	}
	log.Info("Metrics server running", zap.Int("port", port))
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal("unable to start metrics server", zap.String("error", err.Error()), zap.Int("port", port))
	}
}

// Shutdown the server
func (srv *Server) Shutdown() {
	srv.healthServer.Shutdown()
	srv.log.Info("Graceful stop")
	srv.grpcServer.GracefulStop()
}

func annotator(ctx context.Context, req *http.Request) metadata.MD {
	return metadata.New(map[string]string{"requestURI": req.Host + req.URL.RequestURI()})
}

func forwardResponse(ctx context.Context, w http.ResponseWriter, message protobuf.Message) error {
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		return nil
	}

	if vals := md.HeaderMD.Get("x-http-code"); len(vals) > 0 {
		code, err := strconv.Atoi(vals[0])
		if err != nil {
			return err
		}
		w.WriteHeader(code)
		delete(md.HeaderMD, "x-http-code")
		delete(w.Header(), "Grpc-Metadata-X-Http-Code")
	}

	return nil
}

func outcomingHeaderMatcher(header string) (string, bool) {
	switch header {
	case "location", "authorization", "access-control-expose-headers":
		return header, true
	default:
		return header, false
	}
}
