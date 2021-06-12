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

	"github.com/walterwanderley/xo-grpc/_examples/northwind/proto"
	pb_Category "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/category"
	pb_Customer "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/customer"
	pb_CustomerCustomerDemo "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/customer_customer_demo"
	pb_CustomerDemographic "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/customer_demographic"
	pb_Employee "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/employee"
	pb_EmployeeTerritory "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/employee_territory"
	pb_Order "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/order"
	pb_OrderDetail "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/order_detail"
	pb_Product "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/product"
	pb_Region "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/region"
	pb_Shipper "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/shipper"
	pb_Supplier "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/supplier"
	pb_Territory "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/territory"
	pb_UsState "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/us_state"
)

const readTimeout = 15 * time.Second

// Server represents a gRPC server
type Server struct {
	cfg                         Config
	log                         *zap.Logger
	grpcServer                  *grpc.Server
	healthServer                *health.Server
	TerritoryService            pb_Territory.TerritoryServiceServer
	ShipperService              pb_Shipper.ShipperServiceServer
	EmployeeTerritoryService    pb_EmployeeTerritory.EmployeeTerritoryServiceServer
	CustomerService             pb_Customer.CustomerServiceServer
	RegionService               pb_Region.RegionServiceServer
	EmployeeService             pb_Employee.EmployeeServiceServer
	UsStateService              pb_UsState.UsStateServiceServer
	OrderService                pb_Order.OrderServiceServer
	OrderDetailService          pb_OrderDetail.OrderDetailServiceServer
	CustomerDemographicService  pb_CustomerDemographic.CustomerDemographicServiceServer
	CategoryService             pb_Category.CategoryServiceServer
	SupplierService             pb_Supplier.SupplierServiceServer
	CustomerCustomerDemoService pb_CustomerCustomerDemo.CustomerCustomerDemoServiceServer
	ProductService              pb_Product.ProductServiceServer
}

// New gRPC server
func New(cfg Config, log *zap.Logger,
	TerritoryService pb_Territory.TerritoryServiceServer,
	ShipperService pb_Shipper.ShipperServiceServer,
	EmployeeTerritoryService pb_EmployeeTerritory.EmployeeTerritoryServiceServer,
	CustomerService pb_Customer.CustomerServiceServer,
	RegionService pb_Region.RegionServiceServer,
	EmployeeService pb_Employee.EmployeeServiceServer,
	UsStateService pb_UsState.UsStateServiceServer,
	OrderService pb_Order.OrderServiceServer,
	OrderDetailService pb_OrderDetail.OrderDetailServiceServer,
	CustomerDemographicService pb_CustomerDemographic.CustomerDemographicServiceServer,
	CategoryService pb_Category.CategoryServiceServer,
	SupplierService pb_Supplier.SupplierServiceServer,
	CustomerCustomerDemoService pb_CustomerCustomerDemo.CustomerCustomerDemoServiceServer,
	ProductService pb_Product.ProductServiceServer,

) *Server {
	return &Server{
		cfg:                         cfg,
		log:                         log,
		TerritoryService:            TerritoryService,
		ShipperService:              ShipperService,
		EmployeeTerritoryService:    EmployeeTerritoryService,
		CustomerService:             CustomerService,
		RegionService:               RegionService,
		EmployeeService:             EmployeeService,
		UsStateService:              UsStateService,
		OrderService:                OrderService,
		OrderDetailService:          OrderDetailService,
		CustomerDemographicService:  CustomerDemographicService,
		CategoryService:             CategoryService,
		SupplierService:             SupplierService,
		CustomerCustomerDemoService: CustomerCustomerDemoService,
		ProductService:              ProductService,
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

	pb_Territory.RegisterTerritoryServiceServer(srv.grpcServer, srv.TerritoryService)
	pb_Shipper.RegisterShipperServiceServer(srv.grpcServer, srv.ShipperService)
	pb_EmployeeTerritory.RegisterEmployeeTerritoryServiceServer(srv.grpcServer, srv.EmployeeTerritoryService)
	pb_Customer.RegisterCustomerServiceServer(srv.grpcServer, srv.CustomerService)
	pb_Region.RegisterRegionServiceServer(srv.grpcServer, srv.RegionService)
	pb_Employee.RegisterEmployeeServiceServer(srv.grpcServer, srv.EmployeeService)
	pb_UsState.RegisterUsStateServiceServer(srv.grpcServer, srv.UsStateService)
	pb_Order.RegisterOrderServiceServer(srv.grpcServer, srv.OrderService)
	pb_OrderDetail.RegisterOrderDetailServiceServer(srv.grpcServer, srv.OrderDetailService)
	pb_CustomerDemographic.RegisterCustomerDemographicServiceServer(srv.grpcServer, srv.CustomerDemographicService)
	pb_Category.RegisterCategoryServiceServer(srv.grpcServer, srv.CategoryService)
	pb_Supplier.RegisterSupplierServiceServer(srv.grpcServer, srv.SupplierService)
	pb_CustomerCustomerDemo.RegisterCustomerCustomerDemoServiceServer(srv.grpcServer, srv.CustomerCustomerDemoService)
	pb_Product.RegisterProductServiceServer(srv.grpcServer, srv.ProductService)

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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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
	if err := pb_Territory.RegisterTerritoryServiceHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}
	if err := pb_Shipper.RegisterShipperServiceHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}
	if err := pb_EmployeeTerritory.RegisterEmployeeTerritoryServiceHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}
	if err := pb_Customer.RegisterCustomerServiceHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}
	if err := pb_Region.RegisterRegionServiceHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}
	if err := pb_Employee.RegisterEmployeeServiceHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}
	if err := pb_UsState.RegisterUsStateServiceHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}
	if err := pb_Order.RegisterOrderServiceHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}
	if err := pb_OrderDetail.RegisterOrderDetailServiceHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}
	if err := pb_CustomerDemographic.RegisterCustomerDemographicServiceHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}
	if err := pb_Category.RegisterCategoryServiceHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}
	if err := pb_Supplier.RegisterSupplierServiceHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}
	if err := pb_CustomerCustomerDemo.RegisterCustomerCustomerDemoServiceHandler(context.Background(), gwmux, cc); err != nil {
		return err
	}
	if err := pb_Product.RegisterProductServiceHandler(context.Background(), gwmux, cc); err != nil {
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
