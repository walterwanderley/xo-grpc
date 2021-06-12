package application

import (
	"context"
	"database/sql"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	models "github.com/walterwanderley/xo-grpc/_examples/northwind/internal/models"
	pb "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/customer_customer_demo"
	typespb "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/typespb"
)

type CustomerCustomerDemoService struct {
	pb.UnimplementedCustomerCustomerDemoServiceServer
	db *sql.DB
}

func NewCustomerCustomerDemoService(db *sql.DB) *CustomerCustomerDemoService {
	return &CustomerCustomerDemoService{db: db}
}

func (s *CustomerCustomerDemoService) Customer(ctx context.Context, req *pb.CustomerRequest) (res *typespb.Customer, err error) {
	var m models.CustomerCustomerDemo
	m.CustomerID = req.GetCustomerID()

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

func (s *CustomerCustomerDemoService) CustomerCustomerDemoByCustomerIDCustomerTypeID(ctx context.Context, req *pb.CustomerCustomerDemoByCustomerIDCustomerTypeIDRequest) (res *typespb.CustomerCustomerDemo, err error) {

	customerID := req.GetCustomerID()
	customerTypeID := req.GetCustomerTypeID()

	result, err := models.CustomerCustomerDemoByCustomerIDCustomerTypeID(ctx, s.db, customerID, customerTypeID)
	if err != nil {
		return
	}

	res = new(typespb.CustomerCustomerDemo)
	res.CustomerID = result.CustomerID
	res.CustomerTypeID = result.CustomerTypeID

	return
}

func (s *CustomerCustomerDemoService) CustomerDemographic(ctx context.Context, req *pb.CustomerDemographicRequest) (res *typespb.CustomerDemographic, err error) {
	var m models.CustomerCustomerDemo
	m.CustomerTypeID = req.GetCustomerTypeID()

	result, err := m.CustomerDemographic(ctx, s.db)
	if err != nil {
		return
	}

	res = new(typespb.CustomerDemographic)
	res.CustomerTypeID = result.CustomerTypeID
	if result.CustomerDesc.Valid {
		res.CustomerDesc = wrapperspb.String(result.CustomerDesc.String)
	}

	return
}

func (s *CustomerCustomerDemoService) Delete(ctx context.Context, req *pb.DeleteRequest) (res *emptypb.Empty, err error) {
	m, err := models.CustomerCustomerDemoByCustomerIDCustomerTypeID(ctx, s.db, req.CustomerID, req.CustomerTypeID)
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

func (s *CustomerCustomerDemoService) Insert(ctx context.Context, req *pb.InsertRequest) (res *emptypb.Empty, err error) {
	var m models.CustomerCustomerDemo
	m.CustomerID = req.GetCustomerID()
	m.CustomerTypeID = req.GetCustomerTypeID()

	err = m.Insert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}
