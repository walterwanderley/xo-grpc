// Code generated by xo-grpc (https://github.com/walterwanderley/xo-grpc).

package application

import (
	"context"
	"database/sql"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "northwind/api/shipper/v1"
	typespb "northwind/api/typespb/v1"
	models "northwind/internal/models"
)

type ShipperService struct {
	pb.UnimplementedShipperServiceServer
	db     *sql.DB
	logger *zap.Logger
}

func NewShipperService(logger *zap.Logger, db *sql.DB) pb.ShipperServiceServer {
	return &ShipperService{logger: logger, db: db}
}

func (s *ShipperService) Delete(ctx context.Context, req *pb.DeleteRequest) (res *emptypb.Empty, err error) {
	m, err := models.ShipperByShipperID(ctx, s.db, sql.NullInt64{Valid: true, Int64: req.ShipperId.Value})
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

func (s *ShipperService) Insert(ctx context.Context, req *pb.InsertRequest) (res *emptypb.Empty, err error) {
	var m models.Shipper
	if v := req.GetPhone(); v != nil {
		m.Phone = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetShipperId(); v != nil {
		m.ShipperID = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetShipperName(); v != nil {
		m.ShipperName = sql.NullString{Valid: true, String: v.Value}
	}

	err = m.Insert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	err = sendResourceLocation(ctx, fmt.Sprintf("/%v", m.ShipperID))

	return
}

func (s *ShipperService) ShipperByShipperID(ctx context.Context, req *pb.ShipperByShipperIDRequest) (res *typespb.Shipper, err error) {

	var shipperID sql.NullInt64
	if v := req.GetShipperId(); v != nil {
		shipperID = sql.NullInt64{Valid: true, Int64: v.Value}
	}

	result, err := models.ShipperByShipperID(ctx, s.db, shipperID)
	if err != nil {
		return
	}

	res = new(typespb.Shipper)
	if result.ShipperID.Valid {
		res.ShipperId = wrapperspb.Int64(result.ShipperID.Int64)
	}
	if result.ShipperName.Valid {
		res.ShipperName = wrapperspb.String(result.ShipperName.String)
	}
	if result.Phone.Valid {
		res.Phone = wrapperspb.String(result.Phone.String)
	}

	return
}

func (s *ShipperService) Update(ctx context.Context, req *pb.UpdateRequest) (res *emptypb.Empty, err error) {
	m, err := models.ShipperByShipperID(ctx, s.db, sql.NullInt64{Valid: true, Int64: req.ShipperId.Value})
	if err != nil {
		return
	}
	if v := req.GetPhone(); v != nil {
		m.Phone = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetShipperId(); v != nil {
		m.ShipperID = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetShipperName(); v != nil {
		m.ShipperName = sql.NullString{Valid: true, String: v.Value}
	}

	err = m.Update(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}

func (s *ShipperService) Upsert(ctx context.Context, req *pb.UpsertRequest) (res *emptypb.Empty, err error) {
	var m models.Shipper
	if v := req.GetPhone(); v != nil {
		m.Phone = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetShipperId(); v != nil {
		m.ShipperID = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetShipperName(); v != nil {
		m.ShipperName = sql.NullString{Valid: true, String: v.Value}
	}

	err = m.Upsert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}
