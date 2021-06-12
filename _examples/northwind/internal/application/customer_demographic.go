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
	pb "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/customer_demographic"
	typespb "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/typespb"
)

type CustomerDemographicService struct {
	pb.UnimplementedCustomerDemographicServiceServer
	db *sql.DB
}

func NewCustomerDemographicService(db *sql.DB) *CustomerDemographicService {
	return &CustomerDemographicService{db: db}
}

func (s *CustomerDemographicService) Insert(ctx context.Context, req *pb.InsertRequest) (res *emptypb.Empty, err error) {
	var m models.CustomerDemographic
	if v := req.GetCustomerDesc(); v != nil {
		m.CustomerDesc = sql.NullString{Valid: true, String: v.Value}
	}
	m.CustomerTypeID = req.GetCustomerTypeID()

	err = m.Insert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		uri := md.Get("requestURI")
		if len(uri) == 1 {
			err = grpc.SendHeader(ctx, metadata.Pairs(
				"location", fmt.Sprintf("%s/%v", uri[0], m.CustomerTypeID),
				"x-http-code", "201"),
			)
			if err != nil {
				return
			}
		}
	}

	return
}

func (s *CustomerDemographicService) Update(ctx context.Context, req *pb.UpdateRequest) (res *emptypb.Empty, err error) {
	m, err := models.CustomerDemographicByCustomerTypeID(ctx, s.db, req.CustomerTypeID)
	if err != nil {
		return
	}
	if v := req.GetCustomerDesc(); v != nil {
		m.CustomerDesc = sql.NullString{Valid: true, String: v.Value}
	}
	m.CustomerTypeID = req.GetCustomerTypeID()

	err = m.Update(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}

func (s *CustomerDemographicService) Upsert(ctx context.Context, req *pb.UpsertRequest) (res *emptypb.Empty, err error) {
	var m models.CustomerDemographic
	if v := req.GetCustomerDesc(); v != nil {
		m.CustomerDesc = sql.NullString{Valid: true, String: v.Value}
	}
	m.CustomerTypeID = req.GetCustomerTypeID()

	err = m.Upsert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}

func (s *CustomerDemographicService) Delete(ctx context.Context, req *pb.DeleteRequest) (res *emptypb.Empty, err error) {
	m, err := models.CustomerDemographicByCustomerTypeID(ctx, s.db, req.CustomerTypeID)
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

func (s *CustomerDemographicService) CustomerDemographicByCustomerTypeID(ctx context.Context, req *pb.CustomerDemographicByCustomerTypeIDRequest) (res *typespb.CustomerDemographic, err error) {

	customerTypeID := req.GetCustomerTypeID()

	result, err := models.CustomerDemographicByCustomerTypeID(ctx, s.db, customerTypeID)
	if err != nil {
		return
	}

	res = new(typespb.CustomerDemographic)
	res.CustomerTypeID = result.CustomerTypeID
	if result.CustomerDesc.Valid {
		res.CustomerDesc = wrapperspb.String(result.CustomerDesc.String)
	}

	return
}