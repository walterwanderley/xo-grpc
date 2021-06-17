package application

import (
	"context"
	"database/sql"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	models "northwind/internal/models"
	"northwind/internal/validation"
	pb "northwind/proto/order"
	typespb "northwind/proto/typespb"
)

type OrderService struct {
	pb.UnimplementedOrderServer
	db     *sql.DB
	logger *zap.Logger
}

func NewOrderService(logger *zap.Logger, db *sql.DB) *OrderService {
	return &OrderService{logger: logger, db: db}
}

func (s *OrderService) Customer(ctx context.Context, req *pb.CustomerRequest) (res *typespb.Customer, err error) {
	m, err := models.OrderByOrderID(ctx, s.db, int16(req.OrderID))
	if err != nil {
		return
	}

	result, err := m.Customer(ctx, s.db)
	if err != nil {
		return
	}

	res = new(typespb.Customer)
	res.CustomerID = result.CustomerID
	res.CompanyName = result.CompanyName
	if result.ContactName.Valid {
		res.ContactName = wrapperspb.String(result.ContactName.String)
	}
	if result.ContactTitle.Valid {
		res.ContactTitle = wrapperspb.String(result.ContactTitle.String)
	}
	if result.Address.Valid {
		res.Address = wrapperspb.String(result.Address.String)
	}
	if result.City.Valid {
		res.City = wrapperspb.String(result.City.String)
	}
	if result.Region.Valid {
		res.Region = wrapperspb.String(result.Region.String)
	}
	if result.PostalCode.Valid {
		res.PostalCode = wrapperspb.String(result.PostalCode.String)
	}
	if result.Country.Valid {
		res.Country = wrapperspb.String(result.Country.String)
	}
	if result.Phone.Valid {
		res.Phone = wrapperspb.String(result.Phone.String)
	}
	if result.Fax.Valid {
		res.Fax = wrapperspb.String(result.Fax.String)
	}

	return
}

func (s *OrderService) Delete(ctx context.Context, req *pb.DeleteRequest) (res *emptypb.Empty, err error) {
	m, err := models.OrderByOrderID(ctx, s.db, int16(req.OrderID))
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

func (s *OrderService) Employee(ctx context.Context, req *pb.EmployeeRequest) (res *typespb.Employee, err error) {
	m, err := models.OrderByOrderID(ctx, s.db, int16(req.OrderID))
	if err != nil {
		return
	}

	result, err := m.Employee(ctx, s.db)
	if err != nil {
		return
	}

	res = new(typespb.Employee)
	res.EmployeeID = int32(result.EmployeeID)
	res.LastName = result.LastName
	res.FirstName = result.FirstName
	if result.Title.Valid {
		res.Title = wrapperspb.String(result.Title.String)
	}
	if result.TitleOfCourtesy.Valid {
		res.TitleOfCourtesy = wrapperspb.String(result.TitleOfCourtesy.String)
	}
	if result.BirthDate.Valid {
		res.BirthDate = timestamppb.New(result.BirthDate.Time)
	}
	if result.HireDate.Valid {
		res.HireDate = timestamppb.New(result.HireDate.Time)
	}
	if result.Address.Valid {
		res.Address = wrapperspb.String(result.Address.String)
	}
	if result.City.Valid {
		res.City = wrapperspb.String(result.City.String)
	}
	if result.Region.Valid {
		res.Region = wrapperspb.String(result.Region.String)
	}
	if result.PostalCode.Valid {
		res.PostalCode = wrapperspb.String(result.PostalCode.String)
	}
	if result.Country.Valid {
		res.Country = wrapperspb.String(result.Country.String)
	}
	if result.HomePhone.Valid {
		res.HomePhone = wrapperspb.String(result.HomePhone.String)
	}
	if result.Extension.Valid {
		res.Extension = wrapperspb.String(result.Extension.String)
	}
	res.Photo = result.Photo
	if result.Notes.Valid {
		res.Notes = wrapperspb.String(result.Notes.String)
	}
	if result.ReportsTo.Valid {
		res.ReportsTo = wrapperspb.Int64(result.ReportsTo.Int64)
	}
	if result.PhotoPath.Valid {
		res.PhotoPath = wrapperspb.String(result.PhotoPath.String)
	}

	return
}

func (s *OrderService) Insert(ctx context.Context, req *pb.InsertRequest) (res *emptypb.Empty, err error) {
	var m models.Order
	if v := req.GetCustomerID(); v != nil {
		m.CustomerID = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetEmployeeID(); v != nil {
		m.EmployeeID = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetFreight(); v != nil {
		m.Freight = sql.NullFloat64{Valid: true, Float64: v.Value}
	}
	if v := req.GetOrderDate(); v != nil {
		if err = v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid OrderDate: %s%w", err.Error(), validation.ErrUserInput)
			return
		}
		if t := v.AsTime(); !t.IsZero() {
			m.OrderDate.Valid = true
			m.OrderDate.Time = t
		}
	}
	m.OrderID = int16(req.GetOrderID())
	if v := req.GetRequiredDate(); v != nil {
		if err = v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid RequiredDate: %s%w", err.Error(), validation.ErrUserInput)
			return
		}
		if t := v.AsTime(); !t.IsZero() {
			m.RequiredDate.Valid = true
			m.RequiredDate.Time = t
		}
	}
	if v := req.GetShipAddress(); v != nil {
		m.ShipAddress = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetShipCity(); v != nil {
		m.ShipCity = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetShipCountry(); v != nil {
		m.ShipCountry = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetShipName(); v != nil {
		m.ShipName = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetShipPostalCode(); v != nil {
		m.ShipPostalCode = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetShipRegion(); v != nil {
		m.ShipRegion = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetShipVia(); v != nil {
		m.ShipVia = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetShippedDate(); v != nil {
		if err = v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid ShippedDate: %s%w", err.Error(), validation.ErrUserInput)
			return
		}
		if t := v.AsTime(); !t.IsZero() {
			m.ShippedDate.Valid = true
			m.ShippedDate.Time = t
		}
	}

	err = m.Insert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	err = sendResourceLocation(ctx, fmt.Sprintf("/%v", m.OrderID))

	return
}

func (s *OrderService) OrderByOrderID(ctx context.Context, req *pb.OrderByOrderIDRequest) (res *typespb.Order, err error) {

	orderID := int16(req.GetOrderID())

	result, err := models.OrderByOrderID(ctx, s.db, orderID)
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

func (s *OrderService) Shipper(ctx context.Context, req *pb.ShipperRequest) (res *typespb.Shipper, err error) {
	m, err := models.OrderByOrderID(ctx, s.db, int16(req.OrderID))
	if err != nil {
		return
	}

	result, err := m.Shipper(ctx, s.db)
	if err != nil {
		return
	}

	res = new(typespb.Shipper)
	res.ShipperID = int32(result.ShipperID)
	res.CompanyName = result.CompanyName
	if result.Phone.Valid {
		res.Phone = wrapperspb.String(result.Phone.String)
	}

	return
}

func (s *OrderService) Update(ctx context.Context, req *pb.UpdateRequest) (res *emptypb.Empty, err error) {
	m, err := models.OrderByOrderID(ctx, s.db, int16(req.OrderID))
	if err != nil {
		return
	}
	if v := req.GetCustomerID(); v != nil {
		m.CustomerID = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetEmployeeID(); v != nil {
		m.EmployeeID = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetFreight(); v != nil {
		m.Freight = sql.NullFloat64{Valid: true, Float64: v.Value}
	}
	if v := req.GetOrderDate(); v != nil {
		if err = v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid OrderDate: %s%w", err.Error(), validation.ErrUserInput)
			return
		}
		if t := v.AsTime(); !t.IsZero() {
			m.OrderDate.Valid = true
			m.OrderDate.Time = t
		}
	}
	m.OrderID = int16(req.GetOrderID())
	if v := req.GetRequiredDate(); v != nil {
		if err = v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid RequiredDate: %s%w", err.Error(), validation.ErrUserInput)
			return
		}
		if t := v.AsTime(); !t.IsZero() {
			m.RequiredDate.Valid = true
			m.RequiredDate.Time = t
		}
	}
	if v := req.GetShipAddress(); v != nil {
		m.ShipAddress = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetShipCity(); v != nil {
		m.ShipCity = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetShipCountry(); v != nil {
		m.ShipCountry = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetShipName(); v != nil {
		m.ShipName = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetShipPostalCode(); v != nil {
		m.ShipPostalCode = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetShipRegion(); v != nil {
		m.ShipRegion = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetShipVia(); v != nil {
		m.ShipVia = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetShippedDate(); v != nil {
		if err = v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid ShippedDate: %s%w", err.Error(), validation.ErrUserInput)
			return
		}
		if t := v.AsTime(); !t.IsZero() {
			m.ShippedDate.Valid = true
			m.ShippedDate.Time = t
		}
	}

	err = m.Update(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}

func (s *OrderService) Upsert(ctx context.Context, req *pb.UpsertRequest) (res *emptypb.Empty, err error) {
	var m models.Order
	if v := req.GetCustomerID(); v != nil {
		m.CustomerID = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetEmployeeID(); v != nil {
		m.EmployeeID = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetFreight(); v != nil {
		m.Freight = sql.NullFloat64{Valid: true, Float64: v.Value}
	}
	if v := req.GetOrderDate(); v != nil {
		if err = v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid OrderDate: %s%w", err.Error(), validation.ErrUserInput)
			return
		}
		if t := v.AsTime(); !t.IsZero() {
			m.OrderDate.Valid = true
			m.OrderDate.Time = t
		}
	}
	m.OrderID = int16(req.GetOrderID())
	if v := req.GetRequiredDate(); v != nil {
		if err = v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid RequiredDate: %s%w", err.Error(), validation.ErrUserInput)
			return
		}
		if t := v.AsTime(); !t.IsZero() {
			m.RequiredDate.Valid = true
			m.RequiredDate.Time = t
		}
	}
	if v := req.GetShipAddress(); v != nil {
		m.ShipAddress = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetShipCity(); v != nil {
		m.ShipCity = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetShipCountry(); v != nil {
		m.ShipCountry = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetShipName(); v != nil {
		m.ShipName = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetShipPostalCode(); v != nil {
		m.ShipPostalCode = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetShipRegion(); v != nil {
		m.ShipRegion = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetShipVia(); v != nil {
		m.ShipVia = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetShippedDate(); v != nil {
		if err = v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid ShippedDate: %s%w", err.Error(), validation.ErrUserInput)
			return
		}
		if t := v.AsTime(); !t.IsZero() {
			m.ShippedDate.Valid = true
			m.ShippedDate.Time = t
		}
	}

	err = m.Upsert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}
