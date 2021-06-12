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
	pb "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/category"
	typespb "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/typespb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	db *sql.DB
}

func NewCategoryService(db *sql.DB) *CategoryService {
	return &CategoryService{db: db}
}

func (s *CategoryService) Insert(ctx context.Context, req *pb.InsertRequest) (res *emptypb.Empty, err error) {
	var m models.Category
	m.CategoryID = int16(req.GetCategoryID())
	m.CategoryName = req.GetCategoryName()
	if v := req.GetDescription(); v != nil {
		m.Description = sql.NullString{Valid: true, String: v.Value}
	}
	m.Picture = req.GetPicture()

	err = m.Insert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		uri := md.Get("requestURI")
		if len(uri) == 1 {
			err = grpc.SendHeader(ctx, metadata.Pairs(
				"location", fmt.Sprintf("%s/%v", uri[0], m.CategoryID),
				"x-http-code", "201"),
			)
			if err != nil {
				return
			}
		}
	}

	return
}

func (s *CategoryService) Update(ctx context.Context, req *pb.UpdateRequest) (res *emptypb.Empty, err error) {
	m, err := models.CategoryByCategoryID(ctx, s.db, int16(req.CategoryID))
	if err != nil {
		return
	}
	m.CategoryID = int16(req.GetCategoryID())
	m.CategoryName = req.GetCategoryName()
	if v := req.GetDescription(); v != nil {
		m.Description = sql.NullString{Valid: true, String: v.Value}
	}
	m.Picture = req.GetPicture()

	err = m.Update(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}

func (s *CategoryService) Upsert(ctx context.Context, req *pb.UpsertRequest) (res *emptypb.Empty, err error) {
	var m models.Category
	m.CategoryID = int16(req.GetCategoryID())
	m.CategoryName = req.GetCategoryName()
	if v := req.GetDescription(); v != nil {
		m.Description = sql.NullString{Valid: true, String: v.Value}
	}
	m.Picture = req.GetPicture()

	err = m.Upsert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}

func (s *CategoryService) Delete(ctx context.Context, req *pb.DeleteRequest) (res *emptypb.Empty, err error) {
	m, err := models.CategoryByCategoryID(ctx, s.db, int16(req.CategoryID))
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

func (s *CategoryService) CategoryByCategoryID(ctx context.Context, req *pb.CategoryByCategoryIDRequest) (res *typespb.Category, err error) {

	categoryID := int16(req.GetCategoryID())

	result, err := models.CategoryByCategoryID(ctx, s.db, categoryID)
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
