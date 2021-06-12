package application

import (
	"context"
	"database/sql"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	models "github.com/walterwanderley/xo-grpc/_examples/northwind/internal/models"
	pb "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/order_detail"
	typespb "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/typespb"
)

type OrderDetailService struct {
	pb.UnimplementedOrderDetailServiceServer
	db *sql.DB
}

func NewOrderDetailService(db *sql.DB) *OrderDetailService {
	return &OrderDetailService{db: db}
}

func (s *OrderDetailService) Insert(ctx context.Context, req *pb.InsertRequest) (res *emptypb.Empty, err error) {
	var m models.OrderDetail
	m.Discount = req.GetDiscount()
	m.OrderID = int16(req.GetOrderID())
	m.ProductID = int16(req.GetProductID())
	m.Quantity = int16(req.GetQuantity())
	m.UnitPrice = req.GetUnitPrice()

	err = m.Insert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}

func (s *OrderDetailService) Update(ctx context.Context, req *pb.UpdateRequest) (res *emptypb.Empty, err error) {
	m, err := models.OrderDetailByOrderIDProductID(ctx, s.db, int16(req.OrderID), int16(req.ProductID))
	if err != nil {
		return
	}
	m.Discount = req.GetDiscount()
	m.OrderID = int16(req.GetOrderID())
	m.ProductID = int16(req.GetProductID())
	m.Quantity = int16(req.GetQuantity())
	m.UnitPrice = req.GetUnitPrice()

	err = m.Update(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}

func (s *OrderDetailService) Upsert(ctx context.Context, req *pb.UpsertRequest) (res *emptypb.Empty, err error) {
	var m models.OrderDetail
	m.Discount = req.GetDiscount()
	m.OrderID = int16(req.GetOrderID())
	m.ProductID = int16(req.GetProductID())
	m.Quantity = int16(req.GetQuantity())
	m.UnitPrice = req.GetUnitPrice()

	err = m.Upsert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}

func (s *OrderDetailService) Delete(ctx context.Context, req *pb.DeleteRequest) (res *emptypb.Empty, err error) {
	m, err := models.OrderDetailByOrderIDProductID(ctx, s.db, int16(req.OrderID), int16(req.ProductID))
	if err != nil {
		return
	}

	err = m.Delete(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}

func (s *OrderDetailService) OrderDetailByOrderIDProductID(ctx context.Context, req *pb.OrderDetailByOrderIDProductIDRequest) (res *typespb.OrderDetail, err error) {

	orderID := int16(req.GetOrderID())
	productID := int16(req.GetProductID())

	result, err := models.OrderDetailByOrderIDProductID(ctx, s.db, orderID, productID)
	if err != nil {
		return
	}

	res = new(typespb.OrderDetail)
	res.OrderID = int32(result.OrderID)
	res.ProductID = int32(result.ProductID)
	res.UnitPrice = result.UnitPrice
	res.Quantity = int32(result.Quantity)
	res.Discount = result.Discount

	return
}

func (s *OrderDetailService) Order(ctx context.Context, req *pb.OrderRequest) (res *typespb.Order, err error) {
	var m models.OrderDetail
	m.OrderID = int16(req.GetOrderID())

	result, err := m.Order(ctx, s.db)
	if err != nil {
		return
	}

	res = new(typespb.Order)
	res.OrderID = int32(result.OrderID)
	if result.CustomerID.Valid {
		res.CustomerID = wrapperspb.String(result.CustomerID.String)
	}
	if result.EmployeeID.Valid {
		res.EmployeeID = wrapperspb.Int64(result.EmployeeID.Int64)
	}
	if result.OrderDate.Valid {
		res.OrderDate = timestamppb.New(result.OrderDate.Time)
	}
	if result.RequiredDate.Valid {
		res.RequiredDate = timestamppb.New(result.RequiredDate.Time)
	}
	if result.ShippedDate.Valid {
		res.ShippedDate = timestamppb.New(result.ShippedDate.Time)
	}
	if result.ShipVia.Valid {
		res.ShipVia = wrapperspb.Int64(result.ShipVia.Int64)
	}
	if result.Freight.Valid {
		res.Freight = wrapperspb.Double(result.Freight.Float64)
	}
	if result.ShipName.Valid {
		res.ShipName = wrapperspb.String(result.ShipName.String)
	}
	if result.ShipAddress.Valid {
		res.ShipAddress = wrapperspb.String(result.ShipAddress.String)
	}
	if result.ShipCity.Valid {
		res.ShipCity = wrapperspb.String(result.ShipCity.String)
	}
	if result.ShipRegion.Valid {
		res.ShipRegion = wrapperspb.String(result.ShipRegion.String)
	}
	if result.ShipPostalCode.Valid {
		res.ShipPostalCode = wrapperspb.String(result.ShipPostalCode.String)
	}
	if result.ShipCountry.Valid {
		res.ShipCountry = wrapperspb.String(result.ShipCountry.String)
	}

	return
}

func (s *OrderDetailService) Product(ctx context.Context, req *pb.ProductRequest) (res *typespb.Product, err error) {
	var m models.OrderDetail
	m.ProductID = int16(req.GetProductID())

	result, err := m.Product(ctx, s.db)
	if err != nil {
		return
	}

	res = new(typespb.Product)
	res.ProductID = int32(result.ProductID)
	res.ProductName = result.ProductName
	if result.SupplierID.Valid {
		res.SupplierID = wrapperspb.Int64(result.SupplierID.Int64)
	}
	if result.CategoryID.Valid {
		res.CategoryID = wrapperspb.Int64(result.CategoryID.Int64)
	}
	if result.QuantityPerUnit.Valid {
		res.QuantityPerUnit = wrapperspb.String(result.QuantityPerUnit.String)
	}
	if result.UnitPrice.Valid {
		res.UnitPrice = wrapperspb.Double(result.UnitPrice.Float64)
	}
	if result.UnitsInStock.Valid {
		res.UnitsInStock = wrapperspb.Int64(result.UnitsInStock.Int64)
	}
	if result.UnitsOnOrder.Valid {
		res.UnitsOnOrder = wrapperspb.Int64(result.UnitsOnOrder.Int64)
	}
	if result.ReorderLevel.Valid {
		res.ReorderLevel = wrapperspb.Int64(result.ReorderLevel.Int64)
	}
	res.Discontinued = int64(result.Discontinued)

	return
}
