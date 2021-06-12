package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"

	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	// database driver
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/walterwanderley/xo-grpc/_examples/northwind/internal/application"
	"github.com/walterwanderley/xo-grpc/_examples/northwind/internal/models"
	"github.com/walterwanderley/xo-grpc/_examples/northwind/internal/server"
)

const serviceName = "github.com/walterwanderley/xo-grpc/_examples/northwind"

func main() {
	var cfg server.Config
	var dbURL string
	var dev bool
	flag.StringVar(&dbURL, "db", "", "The Database connection URL")
	flag.IntVar(&cfg.Port, "port", 5000, "The server port")
	flag.IntVar(&cfg.PrometheusPort, "prometheusPort", 0, "The metrics server port")
	flag.StringVar(&cfg.JaegerAgent, "jaegerAgent", "", "The Jaeger Tracing agent URL")
	flag.StringVar(&cfg.Cert, "cert", "", "The path to the server certificate file in PEM format")
	flag.StringVar(&cfg.Key, "key", "", "The path to the server private key in PEM format")
	flag.BoolVar(&dev, "dev", false, "Set logger to development mode")
	flag.Parse()

	log := logger(dev)
	defer log.Sync()

	if dev {
		models.SetLogger(func(s string, args ...interface{}) {
			params := make([]string, len(args))
			for i, arg := range args {
				params[i] = fmt.Sprintf("%v", arg)
			}
			log.Debug(fmt.Sprintf("%s; Params -> [%s]", s, strings.Join(params, ", ")))
		})
	}

	if _, err := maxprocs.Set(); err != nil {
		log.Error("startup", zap.Error(err))
		os.Exit(1)
	}
	log.Info("startup", zap.Int("GOMAXPROCS", runtime.GOMAXPROCS(0)))

	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatal("Failed to open DB", zap.String("error", err.Error()))
	}

	srv := server.New(cfg, log,
		application.NewCategoryService(db),
		application.NewCustomerService(db),
		application.NewCustomerCustomerDemoService(db),
		application.NewCustomerDemographicService(db),
		application.NewEmployeeService(db),
		application.NewEmployeeTerritoryService(db),
		application.NewOrderService(db),
		application.NewOrderDetailService(db),
		application.NewProductService(db),
		application.NewRegionService(db),
		application.NewShipperService(db),
		application.NewSupplierService(db),
		application.NewTerritoryService(db),
		application.NewUsStateService(db),
	)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-done
		log.Warn("signal detected...", zap.Any("signal", sig))
		srv.Shutdown()
	}()
	if err := srv.ListenAndServe(); err != nil {
		if err.Error() != "mux: server closed" {
			log.Fatal(err.Error())
		}
	}
}

func logger(dev bool) *zap.Logger {
	var config zap.Config
	if dev {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		config = zap.NewProductionConfig()
	}
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = true
	config.InitialFields = map[string]interface{}{
		"service": serviceName,
	}

	log, err := config.Build()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return log
}
