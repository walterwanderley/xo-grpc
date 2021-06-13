package application

import (
	"context"
	"database/sql"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"

	models "northwind/internal/models"
	pb "northwind/proto/region"
	typespb "northwind/proto/typespb"
)

type RegionService struct {
	pb.UnimplementedRegionServiceServer
	db *sql.DB
}

func NewRegionService(db *sql.DB) *RegionService {
	return &RegionService{db: db}
}

func (s *RegionService) Delete(ctx context.Context, req *pb.DeleteRequest) (res *emptypb.Empty, err error) {
	m, err := models.RegionByRegionID(ctx, s.db, int16(req.RegionID))
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

func (s *RegionService) Insert(ctx context.Context, req *pb.InsertRequest) (res *emptypb.Empty, err error) {
	var m models.Region
	m.RegionDescription = req.GetRegionDescription()
	m.RegionID = int16(req.GetRegionID())

	err = m.Insert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	err = sendResourceLocation(ctx, fmt.Sprintf("/%v", m.RegionID))

	return
}

func (s *RegionService) RegionByRegionID(ctx context.Context, req *pb.RegionByRegionIDRequest) (res *typespb.Region, err error) {

	regionID := int16(req.GetRegionID())

	result, err := models.RegionByRegionID(ctx, s.db, regionID)
	if err != nil {
		return
	}

	res = new(typespb.Region)
	res.RegionID = int32(result.RegionID)
	res.RegionDescription = result.RegionDescription

	return
}

func (s *RegionService) Update(ctx context.Context, req *pb.UpdateRequest) (res *emptypb.Empty, err error) {
	m, err := models.RegionByRegionID(ctx, s.db, int16(req.RegionID))
	if err != nil {
		return
	}
	m.RegionDescription = req.GetRegionDescription()
	m.RegionID = int16(req.GetRegionID())

	err = m.Update(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}

func (s *RegionService) Upsert(ctx context.Context, req *pb.UpsertRequest) (res *emptypb.Empty, err error) {
	var m models.Region
	m.RegionDescription = req.GetRegionDescription()
	m.RegionID = int16(req.GetRegionID())

	err = m.Upsert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}
