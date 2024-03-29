// Code generated by xo-grpc (https://github.com/walterwanderley/xo-grpc).

package application

import (
	"context"
	"database/sql"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "northwind/api/product/v1"
	typespb "northwind/api/typespb/v1"
	models "northwind/internal/models"
)

type ProductService struct {
	pb.UnimplementedProductServiceServer
	db     *sql.DB
	logger *zap.Logger
}

func NewProductService(logger *zap.Logger, db *sql.DB) pb.ProductServiceServer {
	return &ProductService{logger: logger, db: db}
}

func (s *ProductService) Category(ctx context.Context, req *pb.CategoryRequest) (res *typespb.Category, err error) {
	m, err := models.ProductByProductID(ctx, s.db, sql.NullInt64{Valid: true, Int64: req.ProductId.Value})
	if err != nil {
		return
	}

	result, err := m.Category(ctx, s.db)
	if err != nil {
		return
	}

	res = new(typespb.Category)
	if result.CategoryID.Valid {
		res.CategoryId = wrapperspb.Int64(result.CategoryID.Int64)
	}
	if result.CategoryName.Valid {
		res.CategoryName = wrapperspb.String(result.CategoryName.String)
	}
	if result.Description.Valid {
		res.Description = wrapperspb.String(result.Description.String)
	}

	return
}

func (s *ProductService) Delete(ctx context.Context, req *pb.DeleteRequest) (res *emptypb.Empty, err error) {
	m, err := models.ProductByProductID(ctx, s.db, sql.NullInt64{Valid: true, Int64: req.ProductId.Value})
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

func (s *ProductService) Insert(ctx context.Context, req *pb.InsertRequest) (res *emptypb.Empty, err error) {
	var m models.Product
	if v := req.GetCategoryId(); v != nil {
		m.CategoryID = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetPrice(); v != nil {
		m.Price = sql.NullFloat64{Valid: true, Float64: v.Value}
	}
	if v := req.GetProductId(); v != nil {
		m.ProductID = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetProductName(); v != nil {
		m.ProductName = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetSupplierId(); v != nil {
		m.SupplierID = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetUnit(); v != nil {
		m.Unit = sql.NullString{Valid: true, String: v.Value}
	}

	err = m.Insert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	err = sendResourceLocation(ctx, fmt.Sprintf("/%v", m.ProductID))

	return
}

func (s *ProductService) ProductByProductID(ctx context.Context, req *pb.ProductByProductIDRequest) (res *typespb.Product, err error) {

	var productID sql.NullInt64
	if v := req.GetProductId(); v != nil {
		productID = sql.NullInt64{Valid: true, Int64: v.Value}
	}

	result, err := models.ProductByProductID(ctx, s.db, productID)
	if err != nil {
		return
	}

	res = new(typespb.Product)
	if result.ProductID.Valid {
		res.ProductId = wrapperspb.Int64(result.ProductID.Int64)
	}
	if result.ProductName.Valid {
		res.ProductName = wrapperspb.String(result.ProductName.String)
	}
	if result.SupplierID.Valid {
		res.SupplierId = wrapperspb.Int64(result.SupplierID.Int64)
	}
	if result.CategoryID.Valid {
		res.CategoryId = wrapperspb.Int64(result.CategoryID.Int64)
	}
	if result.Unit.Valid {
		res.Unit = wrapperspb.String(result.Unit.String)
	}
	if result.Price.Valid {
		res.Price = wrapperspb.Double(result.Price.Float64)
	}

	return
}

func (s *ProductService) Supplier(ctx context.Context, req *pb.SupplierRequest) (res *typespb.Supplier, err error) {
	m, err := models.ProductByProductID(ctx, s.db, sql.NullInt64{Valid: true, Int64: req.ProductId.Value})
	if err != nil {
		return
	}

	result, err := m.Supplier(ctx, s.db)
	if err != nil {
		return
	}

	res = new(typespb.Supplier)
	if result.SupplierID.Valid {
		res.SupplierId = wrapperspb.Int64(result.SupplierID.Int64)
	}
	if result.SupplierName.Valid {
		res.SupplierName = wrapperspb.String(result.SupplierName.String)
	}
	if result.ContactName.Valid {
		res.ContactName = wrapperspb.String(result.ContactName.String)
	}
	if result.Address.Valid {
		res.Address = wrapperspb.String(result.Address.String)
	}
	if result.City.Valid {
		res.City = wrapperspb.String(result.City.String)
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

	return
}

func (s *ProductService) Update(ctx context.Context, req *pb.UpdateRequest) (res *emptypb.Empty, err error) {
	m, err := models.ProductByProductID(ctx, s.db, sql.NullInt64{Valid: true, Int64: req.ProductId.Value})
	if err != nil {
		return
	}
	if v := req.GetCategoryId(); v != nil {
		m.CategoryID = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetPrice(); v != nil {
		m.Price = sql.NullFloat64{Valid: true, Float64: v.Value}
	}
	if v := req.GetProductId(); v != nil {
		m.ProductID = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetProductName(); v != nil {
		m.ProductName = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetSupplierId(); v != nil {
		m.SupplierID = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetUnit(); v != nil {
		m.Unit = sql.NullString{Valid: true, String: v.Value}
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
	if v := req.GetCategoryId(); v != nil {
		m.CategoryID = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetPrice(); v != nil {
		m.Price = sql.NullFloat64{Valid: true, Float64: v.Value}
	}
	if v := req.GetProductId(); v != nil {
		m.ProductID = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetProductName(); v != nil {
		m.ProductName = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetSupplierId(); v != nil {
		m.SupplierID = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetUnit(); v != nil {
		m.Unit = sql.NullString{Valid: true, String: v.Value}
	}

	err = m.Upsert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}
