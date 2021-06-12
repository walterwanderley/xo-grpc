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
	typespb "northwind/proto/typespb"
	pb "northwind/proto/us_state"
)

type UsStateService struct {
	pb.UnimplementedUsStateServiceServer
	db *sql.DB
}

func NewUsStateService(db *sql.DB) *UsStateService {
	return &UsStateService{db: db}
}

func (s *UsStateService) Delete(ctx context.Context, req *pb.DeleteRequest) (res *emptypb.Empty, err error) {
	m, err := models.UsStateByStateID(ctx, s.db, int16(req.StateID))
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

func (s *UsStateService) Insert(ctx context.Context, req *pb.InsertRequest) (res *emptypb.Empty, err error) {
	var m models.UsState
	if v := req.GetStateAbbr(); v != nil {
		m.StateAbbr = sql.NullString{Valid: true, String: v.Value}
	}
	m.StateID = int16(req.GetStateID())
	if v := req.GetStateName(); v != nil {
		m.StateName = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetStateRegion(); v != nil {
		m.StateRegion = sql.NullString{Valid: true, String: v.Value}
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
				"location", fmt.Sprintf("%s/%v", uri[0], m.StateID),
				"x-http-code", "201"),
			)
			if err != nil {
				return
			}
		}
	}

	return
}

func (s *UsStateService) Update(ctx context.Context, req *pb.UpdateRequest) (res *emptypb.Empty, err error) {
	m, err := models.UsStateByStateID(ctx, s.db, int16(req.StateID))
	if err != nil {
		return
	}
	if v := req.GetStateAbbr(); v != nil {
		m.StateAbbr = sql.NullString{Valid: true, String: v.Value}
	}
	m.StateID = int16(req.GetStateID())
	if v := req.GetStateName(); v != nil {
		m.StateName = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetStateRegion(); v != nil {
		m.StateRegion = sql.NullString{Valid: true, String: v.Value}
	}

	err = m.Update(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}

func (s *UsStateService) Upsert(ctx context.Context, req *pb.UpsertRequest) (res *emptypb.Empty, err error) {
	var m models.UsState
	if v := req.GetStateAbbr(); v != nil {
		m.StateAbbr = sql.NullString{Valid: true, String: v.Value}
	}
	m.StateID = int16(req.GetStateID())
	if v := req.GetStateName(); v != nil {
		m.StateName = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetStateRegion(); v != nil {
		m.StateRegion = sql.NullString{Valid: true, String: v.Value}
	}

	err = m.Upsert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}

func (s *UsStateService) UsStateByStateID(ctx context.Context, req *pb.UsStateByStateIDRequest) (res *typespb.UsState, err error) {

	stateID := int16(req.GetStateID())

	result, err := models.UsStateByStateID(ctx, s.db, stateID)
	if err != nil {
		return
	}

	res = new(typespb.UsState)
	res.StateID = int32(result.StateID)
	if result.StateName.Valid {
		res.StateName = wrapperspb.String(result.StateName.String)
	}
	if result.StateAbbr.Valid {
		res.StateAbbr = wrapperspb.String(result.StateAbbr.String)
	}
	if result.StateRegion.Valid {
		res.StateRegion = wrapperspb.String(result.StateRegion.String)
	}

	return
}
