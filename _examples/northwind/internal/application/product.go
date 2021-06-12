package application

import (
	"context"
	"database/sql"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	models "github.com/walterwanderley/xo-grpc/_examples/northwind/internal/models"
	pb "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/product"
	typespb "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/typespb"
)

type ProductService struct {
	pb.UnimplementedProductServiceServer
	db *sql.DB
}

func NewProductService(db *sql.DB) *ProductService {
	return &ProductService{db: db}
}

func (s *ProductService) Insert(ctx context.Context, req *pb.InsertRequest) (res *emptypb.Empty, err error) {
	var m models.Product
	if v := req.GetCategoryID(); v != nil {
		m.CategoryID = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	m.Discontinued = int(req.GetDiscontinued())
	m.ProductID = int16(req.GetProductID())
	m.ProductName = req.GetProductName()
	if v := req.GetQuantityPerUnit(); v != nil {
		m.QuantityPerUnit = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetReorderLevel(); v != nil {
		m.ReorderLevel = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetSupplierID(); v != nil {
		m.SupplierID = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetUnitPrice(); v != nil {
		m.UnitPrice = sql.NullFloat64{Valid: true, Float64: v.Value}
	}
	if v := req.GetUnitsInStock(); v != nil {
		m.UnitsInStock = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetUnitsOnOrder(); v != nil {
		m.UnitsOnOrder = sql.NullInt64{Valid: true, Int64: v.Value}
	}

	err = m.Insert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		uri := md.Get("requestURI")
		if len(uri) == 1 {
			err = grpc.SendHeader(ctx, metadata.Pairs(
				"location", fmt.Sprintf("%s/%v", uri[0], m.ProductID),
				"x-http-code", "201"),
			)
			if err != nil {
				return
			}
		}
	}

	return
}

func (s *ProductService) Update(ctx context.Context, req *pb.UpdateRequest) (res *emptypb.Empty, err error) {
	m, err := models.ProductByProductID(ctx, s.db, int16(req.ProductID))
	if err != nil {
		return
	}
	if v := req.GetCategoryID(); v != nil {
		m.CategoryID = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	m.Discontinued = int(req.GetDiscontinued())
	m.ProductID = int16(req.GetProductID())
	m.ProductName = req.GetProductName()
	if v := req.GetQuantityPerUnit(); v != nil {
		m.QuantityPerUnit = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetReorderLevel(); v != nil {
		m.ReorderLevel = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetSupplierID(); v != nil {
		m.SupplierID = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetUnitPrice(); v != nil {
		m.UnitPrice = sql.NullFloat64{Valid: true, Float64: v.Value}
	}
	if v := req.GetUnitsInStock(); v != nil {
		m.UnitsInStock = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetUnitsOnOrder(); v != nil {
		m.UnitsOnOrder = sql.NullInt64{Valid: true, Int64: v.Value}
	}

	err = m.Update(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}

func (s *ProductService) Upsert(ctx context.Context, req *pb.UpsertRequest) (res *emptypb.Empty, err error) {
	var m models.Product
	if v := req.GetCategoryID(); v != nil {
		m.CategoryID = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	m.Discontinued = int(req.GetDiscontinued())
	m.ProductID = int16(req.GetProductID())
	m.ProductName = req.GetProductName()
	if v := req.GetQuantityPerUnit(); v != nil {
		m.QuantityPerUnit = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetReorderLevel(); v != nil {
		m.ReorderLevel = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetSupplierID(); v != nil {
		m.SupplierID = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetUnitPrice(); v != nil {
		m.UnitPrice = sql.NullFloat64{Valid: true, Float64: v.Value}
	}
	if v := req.GetUnitsInStock(); v != nil {
		m.UnitsInStock = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetUnitsOnOrder(); v != nil {
		m.UnitsOnOrder = sql.NullInt64{Valid: true, Int64: v.Value}
	}

	err = m.Upsert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}

func (s *ProductService) Delete(ctx context.Context, req *pb.DeleteRequest) (res *emptypb.Empty, err error) {
	m, err := models.ProductByProductID(ctx, s.db, int16(req.ProductID))
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

func (s *ProductService) ProductByProductID(ctx context.Context, req *pb.ProductByProductIDRequest) (res *typespb.Product, err error) {

	productID := int16(req.GetProductID())

	result, err := models.ProductByProductID(ctx, s.db, productID)
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

func (s *ProductService) Category(ctx context.Context, req *pb.CategoryRequest) (res *typespb.Category, err error) {
	var m models.Product
	if v := req.GetCategoryID(); v != nil {
		m.CategoryID = sql.NullInt64{Valid: true, Int64: v.Value}
	}

	result, err := m.Category(ctx, s.db)
	if err != nil {
		return
	}

	res = new(typespb.Category)
	res.CategoryID = int32(result.CategoryID)
	res.CategoryName = result.CategoryName
	if result.Description.Valid {
		res.Description = wrapperspb.String(result.Description.String)
	}
	res.Picture = result.Picture

	return
}

func (s *ProductService) Supplier(ctx context.Context, req *pb.SupplierRequest) (res *typespb.Supplier, err error) {
	var m models.Product
	if v := req.GetSupplierID(); v != nil {
		m.SupplierID = sql.NullInt64{Valid: true, Int64: v.Value}
	}

	result, err := m.Supplier(ctx, s.db)
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
