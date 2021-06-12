package application

import (
	"context"
	"database/sql"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"

	models "northwind/internal/models"
	pb "northwind/proto/territory"
	typespb "northwind/proto/typespb"
)

type TerritoryService struct {
	pb.UnimplementedTerritoryServiceServer
	db *sql.DB
}

func NewTerritoryService(db *sql.DB) *TerritoryService {
	return &TerritoryService{db: db}
}

func (s *TerritoryService) Delete(ctx context.Context, req *pb.DeleteRequest) (res *emptypb.Empty, err error) {
	m, err := models.TerritoryByTerritoryID(ctx, s.db, req.TerritoryID)
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

func (s *TerritoryService) Insert(ctx context.Context, req *pb.InsertRequest) (res *emptypb.Empty, err error) {
	var m models.Territory
	m.RegionID = int16(req.GetRegionID())
	m.TerritoryDescription = req.GetTerritoryDescription()
	m.TerritoryID = req.GetTerritoryID()

	err = m.Insert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		uri := md.Get("requestURI")
		if len(uri) == 1 {
			err = grpc.SendHeader(ctx, metadata.Pairs(
				"location", fmt.Sprintf("%s/%v", uri[0], m.TerritoryID),
				"x-http-code", "201"),
			)
			if err != nil {
				return
			}
		}
	}

	return
}

func (s *TerritoryService) Region(ctx context.Context, req *pb.RegionRequest) (res *typespb.Region, err error) {
	var m models.Territory
	m.RegionID = int16(req.GetRegionID())

	result, err := m.Region(ctx, s.db)
	if err != nil {
		return
	}

	res = new(typespb.Region)
	res.RegionID = int32(result.RegionID)
	res.RegionDescription = result.RegionDescription

	return
}

func (s *TerritoryService) TerritoryByTerritoryID(ctx context.Context, req *pb.TerritoryByTerritoryIDRequest) (res *typespb.Territory, err error) {

	territoryID := req.GetTerritoryID()

	result, err := models.TerritoryByTerritoryID(ctx, s.db, territoryID)
	if err != nil {
		return
	}

	res = new(typespb.Territory)
	res.TerritoryID = result.TerritoryID
	res.TerritoryDescription = result.TerritoryDescription
	res.RegionID = int32(result.RegionID)

	return
}

func (s *TerritoryService) Update(ctx context.Context, req *pb.UpdateRequest) (res *emptypb.Empty, err error) {
	m, err := models.TerritoryByTerritoryID(ctx, s.db, req.TerritoryID)
	if err != nil {
		return
	}
	m.RegionID = int16(req.GetRegionID())
	m.TerritoryDescription = req.GetTerritoryDescription()
	m.TerritoryID = req.GetTerritoryID()

	err = m.Update(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}

func (s *TerritoryService) Upsert(ctx context.Context, req *pb.UpsertRequest) (res *emptypb.Empty, err error) {
	var m models.Territory
	m.RegionID = int16(req.GetRegionID())
	m.TerritoryDescription = req.GetTerritoryDescription()
	m.TerritoryID = req.GetTerritoryID()

	err = m.Upsert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}
