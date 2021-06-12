package application

import (
	"context"
	"database/sql"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	models "northwind/internal/models"
	pb "northwind/proto/supplier"
	typespb "northwind/proto/typespb"
)

type SupplierService struct {
	pb.UnimplementedSupplierServiceServer
	db *sql.DB
}

func NewSupplierService(db *sql.DB) *SupplierService {
	return &SupplierService{db: db}
}

func (s *SupplierService) Delete(ctx context.Context, req *pb.DeleteRequest) (res *emptypb.Empty, err error) {
	m, err := models.SupplierBySupplierID(ctx, s.db, int16(req.SupplierID))
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

func (s *SupplierService) Insert(ctx context.Context, req *pb.InsertRequest) (res *emptypb.Empty, err error) {
	var m models.Supplier
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
	if v := req.GetFax(); v != nil {
		m.Fax = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetHomepage(); v != nil {
		m.Homepage = sql.NullString{Valid: true, String: v.Value}
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
	m.SupplierID = int16(req.GetSupplierID())

	err = m.Insert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		uri := md.Get("requestURI")
		if len(uri) == 1 {
			err = grpc.SendHeader(ctx, metadata.Pairs(
				"location", fmt.Sprintf("%s/%v", uri[0], m.SupplierID),
				"x-http-code", "201"),
			)
			if err != nil {
				return
			}
		}
	}

	return
}

func (s *SupplierService) SupplierBySupplierID(ctx context.Context, req *pb.SupplierBySupplierIDRequest) (res *typespb.Supplier, err error) {

	supplierID := int16(req.GetSupplierID())

	result, err := models.SupplierBySupplierID(ctx, s.db, supplierID)
	if err != nil {
		return
	}

	res = new(typespb.Supplier)
	res.SupplierID = int32(result.SupplierID)
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
	if result.Homepage.Valid {
		res.Homepage = wrapperspb.String(result.Homepage.String)
	}

	return
}

func (s *SupplierService) Update(ctx context.Context, req *pb.UpdateRequest) (res *emptypb.Empty, err error) {
	m, err := models.SupplierBySupplierID(ctx, s.db, int16(req.SupplierID))
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
	if v := req.GetFax(); v != nil {
		m.Fax = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetHomepage(); v != nil {
		m.Homepage = sql.NullString{Valid: true, String: v.Value}
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
	m.SupplierID = int16(req.GetSupplierID())

	err = m.Update(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}

func (s *SupplierService) Upsert(ctx context.Context, req *pb.UpsertRequest) (res *emptypb.Empty, err error) {
	var m models.Supplier
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
	if v := req.GetFax(); v != nil {
		m.Fax = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetHomepage(); v != nil {
		m.Homepage = sql.NullString{Valid: true, String: v.Value}
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
	m.SupplierID = int16(req.GetSupplierID())

	err = m.Upsert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}
