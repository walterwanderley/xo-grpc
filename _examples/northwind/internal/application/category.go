package application

import (
	"context"
	"database/sql"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	models "northwind/internal/models"
	pb "northwind/proto/category"
	typespb "northwind/proto/typespb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	db *sql.DB
}

func NewCategoryService(db *sql.DB) *CategoryService {
	return &CategoryService{db: db}
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

	err = sendResourceLocation(ctx, fmt.Sprintf("/%v", m.CategoryID))

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
