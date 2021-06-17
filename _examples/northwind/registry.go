package main

import (
	"database/sql"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"northwind/internal/application"
	"northwind/internal/server"
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

func registerServer(logger *zap.Logger, db *sql.DB) server.RegisterServer {
	return func(grpcServer *grpc.Server) {
		pb_Category.RegisterCategoryServer(grpcServer, application.NewCategoryService(logger, db))
		pb_Customer.RegisterCustomerServer(grpcServer, application.NewCustomerService(logger, db))
		pb_CustomerCustomerDemo.RegisterCustomerCustomerDemoServer(grpcServer, application.NewCustomerCustomerDemoService(logger, db))
		pb_CustomerDemographic.RegisterCustomerDemographicServer(grpcServer, application.NewCustomerDemographicService(logger, db))
		pb_Employee.RegisterEmployeeServer(grpcServer, application.NewEmployeeService(logger, db))
		pb_EmployeeTerritory.RegisterEmployeeTerritoryServer(grpcServer, application.NewEmployeeTerritoryService(logger, db))
		pb_Order.RegisterOrderServer(grpcServer, application.NewOrderService(logger, db))
		pb_OrderDetail.RegisterOrderDetailServer(grpcServer, application.NewOrderDetailService(logger, db))
		pb_Product.RegisterProductServer(grpcServer, application.NewProductService(logger, db))
		pb_Region.RegisterRegionServer(grpcServer, application.NewRegionService(logger, db))
		pb_Shipper.RegisterShipperServer(grpcServer, application.NewShipperService(logger, db))
		pb_Supplier.RegisterSupplierServer(grpcServer, application.NewSupplierService(logger, db))
		pb_Territory.RegisterTerritoryServer(grpcServer, application.NewTerritoryService(logger, db))
		pb_UsState.RegisterUsStateServer(grpcServer, application.NewUsStateService(logger, db))

	}
}

func registerHandlers() []server.RegisterHandler {
	var handlers []server.RegisterHandler

	handlers = append(handlers, pb_Category.RegisterCategoryHandler)
	handlers = append(handlers, pb_Customer.RegisterCustomerHandler)
	handlers = append(handlers, pb_CustomerCustomerDemo.RegisterCustomerCustomerDemoHandler)
	handlers = append(handlers, pb_CustomerDemographic.RegisterCustomerDemographicHandler)
	handlers = append(handlers, pb_Employee.RegisterEmployeeHandler)
	handlers = append(handlers, pb_EmployeeTerritory.RegisterEmployeeTerritoryHandler)
	handlers = append(handlers, pb_Order.RegisterOrderHandler)
	handlers = append(handlers, pb_OrderDetail.RegisterOrderDetailHandler)
	handlers = append(handlers, pb_Product.RegisterProductHandler)
	handlers = append(handlers, pb_Region.RegisterRegionHandler)
	handlers = append(handlers, pb_Shipper.RegisterShipperHandler)
	handlers = append(handlers, pb_Supplier.RegisterSupplierHandler)
	handlers = append(handlers, pb_Territory.RegisterTerritoryHandler)
	handlers = append(handlers, pb_UsState.RegisterUsStateHandler)

	return handlers
}
