package application

import (
	"context"
	"database/sql"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	models "northwind/internal/models"
	pb "northwind/proto/customer"
	typespb "northwind/proto/typespb"
)

type CustomerService struct {
	pb.UnimplementedCustomerServiceServer
	db *sql.DB
}

func NewCustomerService(db *sql.DB) *CustomerService {
	return &CustomerService{db: db}
}

func (s *CustomerService) CustomerByCustomerID(ctx context.Context, req *pb.CustomerByCustomerIDRequest) (res *typespb.Customer, err error) {

	customerID := req.GetCustomerID()

	result, err := models.CustomerByCustomerID(ctx, s.db, customerID)
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

func (s *CustomerService) Delete(ctx context.Context, req *pb.DeleteRequest) (res *emptypb.Empty, err error) {
	m, err := models.CustomerByCustomerID(ctx, s.db, req.CustomerID)
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

func (s *CustomerService) Insert(ctx context.Context, req *pb.InsertRequest) (res *emptypb.Empty, err error) {
	var m models.Customer
	if v := req.GetAddress(); v != nil {
		m.Address = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetCity(); v != nil {
		m.City = sql.NullString{Valid: true, String: v.Value}
	}
	m.CompanyName = req.GetCompanyName()
	if v := req.GetContactName(); v != nil {
		m.ContactName = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetContactTitle(); v != nil {
		m.ContactTitle = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetCountry(); v != nil {
		m.Country = sql.NullString{Valid: true, String: v.Value}
	}
	m.CustomerID = req.GetCustomerID()
	if v := req.GetFax(); v != nil {
		m.Fax = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetPhone(); v != nil {
		m.Phone = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetPostalCode(); v != nil {
		m.PostalCode = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetRegion(); v != nil {
		m.Region = sql.NullString{Valid: true, String: v.Value}
	}

	err = m.Insert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	err = sendResourceLocation(ctx, fmt.Sprintf("/%v", m.CustomerID))

	return
}

func (s *CustomerService) Update(ctx context.Context, req *pb.UpdateRequest) (res *emptypb.Empty, err error) {
	m, err := models.CustomerByCustomerID(ctx, s.db, req.CustomerID)
	if err != nil {
		return
	}
	if v := req.GetAddress(); v != nil {
		m.Address = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetCity(); v != nil {
		m.City = sql.NullString{Valid: true, String: v.Value}
	}
	m.CompanyName = req.GetCompanyName()
	if v := req.GetContactName(); v != nil {
		m.ContactName = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetContactTitle(); v != nil {
		m.ContactTitle = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetCountry(); v != nil {
		m.Country = sql.NullString{Valid: true, String: v.Value}
	}
	m.CustomerID = req.GetCustomerID()
	if v := req.GetFax(); v != nil {
		m.Fax = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetPhone(); v != nil {
		m.Phone = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetPostalCode(); v != nil {
		m.PostalCode = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetRegion(); v != nil {
		m.Region = sql.NullString{Valid: true, String: v.Value}
	}

	err = m.Update(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}

func (s *CustomerService) Upsert(ctx context.Context, req *pb.UpsertRequest) (res *emptypb.Empty, err error) {
	var m models.Customer
	if v := req.GetAddress(); v != nil {
		m.Address = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetCity(); v != nil {
		m.City = sql.NullString{Valid: true, String: v.Value}
	}
	m.CompanyName = req.GetCompanyName()
	if v := req.GetContactName(); v != nil {
		m.ContactName = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetContactTitle(); v != nil {
		m.ContactTitle = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetCountry(); v != nil {
		m.Country = sql.NullString{Valid: true, String: v.Value}
	}
	m.CustomerID = req.GetCustomerID()
	if v := req.GetFax(); v != nil {
		m.Fax = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetPhone(); v != nil {
		m.Phone = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetPostalCode(); v != nil {
		m.PostalCode = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetRegion(); v != nil {
		m.Region = sql.NullString{Valid: true, String: v.Value}
	}

	err = m.Upsert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}
