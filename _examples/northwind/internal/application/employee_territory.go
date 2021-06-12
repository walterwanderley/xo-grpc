package application

import (
	"context"
	"database/sql"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	models "github.com/walterwanderley/xo-grpc/_examples/northwind/internal/models"
	pb "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/employee_territory"
	typespb "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/typespb"
)

type EmployeeTerritoryService struct {
	pb.UnimplementedEmployeeTerritoryServiceServer
	db *sql.DB
}

func NewEmployeeTerritoryService(db *sql.DB) *EmployeeTerritoryService {
	return &EmployeeTerritoryService{db: db}
}

func (s *EmployeeTerritoryService) Delete(ctx context.Context, req *pb.DeleteRequest) (res *emptypb.Empty, err error) {
	m, err := models.EmployeeTerritoryByEmployeeIDTerritoryID(ctx, s.db, int16(req.EmployeeID), req.TerritoryID)
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

func (s *EmployeeTerritoryService) Employee(ctx context.Context, req *pb.EmployeeRequest) (res *typespb.Employee, err error) {
	var m models.EmployeeTerritory
	m.EmployeeID = int16(req.GetEmployeeID())

	result, err := m.Employee(ctx, s.db)
	if err != nil {
		return
	}

	res = new(typespb.Employee)
	res.EmployeeID = int32(result.EmployeeID)
	res.LastName = result.LastName
	res.FirstName = result.FirstName
	if result.Title.Valid {
		res.Title = wrapperspb.String(result.Title.String)
	}
	if result.TitleOfCourtesy.Valid {
		res.TitleOfCourtesy = wrapperspb.String(result.TitleOfCourtesy.String)
	}
	if result.BirthDate.Valid {
		res.BirthDate = timestamppb.New(result.BirthDate.Time)
	}
	if result.HireDate.Valid {
		res.HireDate = timestamppb.New(result.HireDate.Time)
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
	if result.HomePhone.Valid {
		res.HomePhone = wrapperspb.String(result.HomePhone.String)
	}
	if result.Extension.Valid {
		res.Extension = wrapperspb.String(result.Extension.String)
	}
	res.Photo = result.Photo
	if result.Notes.Valid {
		res.Notes = wrapperspb.String(result.Notes.String)
	}
	if result.ReportsTo.Valid {
		res.ReportsTo = wrapperspb.Int64(result.ReportsTo.Int64)
	}
	if result.PhotoPath.Valid {
		res.PhotoPath = wrapperspb.String(result.PhotoPath.String)
	}

	return
}

func (s *EmployeeTerritoryService) EmployeeTerritoryByEmployeeIDTerritoryID(ctx context.Context, req *pb.EmployeeTerritoryByEmployeeIDTerritoryIDRequest) (res *typespb.EmployeeTerritory, err error) {

	employeeID := int16(req.GetEmployeeID())
	territoryID := req.GetTerritoryID()

	result, err := models.EmployeeTerritoryByEmployeeIDTerritoryID(ctx, s.db, employeeID, territoryID)
	if err != nil {
		return
	}

	res = new(typespb.EmployeeTerritory)
	res.EmployeeID = int32(result.EmployeeID)
	res.TerritoryID = result.TerritoryID

	return
}

func (s *EmployeeTerritoryService) Insert(ctx context.Context, req *pb.InsertRequest) (res *emptypb.Empty, err error) {
	var m models.EmployeeTerritory
	m.EmployeeID = int16(req.GetEmployeeID())
	m.TerritoryID = req.GetTerritoryID()

	err = m.Insert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}

func (s *EmployeeTerritoryService) Territory(ctx context.Context, req *pb.TerritoryRequest) (res *typespb.Territory, err error) {
	var m models.EmployeeTerritory
	m.TerritoryID = req.GetTerritoryID()

	result, err := m.Territory(ctx, s.db)
	if err != nil {
		return
	}

	res = new(typespb.Territory)
	res.TerritoryID = result.TerritoryID
	res.TerritoryDescription = result.TerritoryDescription
	res.RegionID = int32(result.RegionID)

	return
}
