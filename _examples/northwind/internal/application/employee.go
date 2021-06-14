package application

import (
	"context"
	"database/sql"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	models "northwind/internal/models"
	"northwind/internal/validation"
	pb "northwind/proto/employee"
	typespb "northwind/proto/typespb"
)

type EmployeeService struct {
	pb.UnimplementedEmployeeServer
	db *sql.DB
}

func NewEmployeeService(db *sql.DB) *EmployeeService {
	return &EmployeeService{db: db}
}

func (s *EmployeeService) Delete(ctx context.Context, req *pb.DeleteRequest) (res *emptypb.Empty, err error) {
	m, err := models.EmployeeByEmployeeID(ctx, s.db, int16(req.EmployeeID))
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

func (s *EmployeeService) Employee(ctx context.Context, req *pb.EmployeeRequest) (res *typespb.Employee, err error) {
	m, err := models.EmployeeByEmployeeID(ctx, s.db, int16(req.EmployeeID))
	if err != nil {
		return
	}

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

func (s *EmployeeService) EmployeeByEmployeeID(ctx context.Context, req *pb.EmployeeByEmployeeIDRequest) (res *typespb.Employee, err error) {

	employeeID := int16(req.GetEmployeeID())

	result, err := models.EmployeeByEmployeeID(ctx, s.db, employeeID)
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

func (s *EmployeeService) Insert(ctx context.Context, req *pb.InsertRequest) (res *emptypb.Empty, err error) {
	var m models.Employee
	if v := req.GetAddress(); v != nil {
		m.Address = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetBirthDate(); v != nil {
		if err = v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid BirthDate: %s%w", err.Error(), validation.ErrUserInput)
			return
		}
		if t := v.AsTime(); !t.IsZero() {
			m.BirthDate.Valid = true
			m.BirthDate.Time = t
		}
	}
	if v := req.GetCity(); v != nil {
		m.City = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetCountry(); v != nil {
		m.Country = sql.NullString{Valid: true, String: v.Value}
	}
	m.EmployeeID = int16(req.GetEmployeeID())
	if v := req.GetExtension(); v != nil {
		m.Extension = sql.NullString{Valid: true, String: v.Value}
	}
	m.FirstName = req.GetFirstName()
	if v := req.GetHireDate(); v != nil {
		if err = v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid HireDate: %s%w", err.Error(), validation.ErrUserInput)
			return
		}
		if t := v.AsTime(); !t.IsZero() {
			m.HireDate.Valid = true
			m.HireDate.Time = t
		}
	}
	if v := req.GetHomePhone(); v != nil {
		m.HomePhone = sql.NullString{Valid: true, String: v.Value}
	}
	m.LastName = req.GetLastName()
	if v := req.GetNotes(); v != nil {
		m.Notes = sql.NullString{Valid: true, String: v.Value}
	}
	m.Photo = req.GetPhoto()
	if v := req.GetPhotoPath(); v != nil {
		m.PhotoPath = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetPostalCode(); v != nil {
		m.PostalCode = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetRegion(); v != nil {
		m.Region = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetReportsTo(); v != nil {
		m.ReportsTo = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetTitle(); v != nil {
		m.Title = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetTitleOfCourtesy(); v != nil {
		m.TitleOfCourtesy = sql.NullString{Valid: true, String: v.Value}
	}

	err = m.Insert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	err = sendResourceLocation(ctx, fmt.Sprintf("/%v", m.EmployeeID))
	return
}

func (s *EmployeeService) Update(ctx context.Context, req *pb.UpdateRequest) (res *emptypb.Empty, err error) {
	m, err := models.EmployeeByEmployeeID(ctx, s.db, int16(req.EmployeeID))
	if err != nil {
		return
	}
	if v := req.GetAddress(); v != nil {
		m.Address = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetBirthDate(); v != nil {
		if err = v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid BirthDate: %s%w", err.Error(), validation.ErrUserInput)
			return
		}
		if t := v.AsTime(); !t.IsZero() {
			m.BirthDate.Valid = true
			m.BirthDate.Time = t
		}
	}
	if v := req.GetCity(); v != nil {
		m.City = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetCountry(); v != nil {
		m.Country = sql.NullString{Valid: true, String: v.Value}
	}
	m.EmployeeID = int16(req.GetEmployeeID())
	if v := req.GetExtension(); v != nil {
		m.Extension = sql.NullString{Valid: true, String: v.Value}
	}
	m.FirstName = req.GetFirstName()
	if v := req.GetHireDate(); v != nil {
		if err = v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid HireDate: %s%w", err.Error(), validation.ErrUserInput)
			return
		}
		if t := v.AsTime(); !t.IsZero() {
			m.HireDate.Valid = true
			m.HireDate.Time = t
		}
	}
	if v := req.GetHomePhone(); v != nil {
		m.HomePhone = sql.NullString{Valid: true, String: v.Value}
	}
	m.LastName = req.GetLastName()
	if v := req.GetNotes(); v != nil {
		m.Notes = sql.NullString{Valid: true, String: v.Value}
	}
	m.Photo = req.GetPhoto()
	if v := req.GetPhotoPath(); v != nil {
		m.PhotoPath = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetPostalCode(); v != nil {
		m.PostalCode = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetRegion(); v != nil {
		m.Region = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetReportsTo(); v != nil {
		m.ReportsTo = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetTitle(); v != nil {
		m.Title = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetTitleOfCourtesy(); v != nil {
		m.TitleOfCourtesy = sql.NullString{Valid: true, String: v.Value}
	}

	err = m.Update(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}

func (s *EmployeeService) Upsert(ctx context.Context, req *pb.UpsertRequest) (res *emptypb.Empty, err error) {
	var m models.Employee
	if v := req.GetAddress(); v != nil {
		m.Address = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetBirthDate(); v != nil {
		if err = v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid BirthDate: %s%w", err.Error(), validation.ErrUserInput)
			return
		}
		if t := v.AsTime(); !t.IsZero() {
			m.BirthDate.Valid = true
			m.BirthDate.Time = t
		}
	}
	if v := req.GetCity(); v != nil {
		m.City = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetCountry(); v != nil {
		m.Country = sql.NullString{Valid: true, String: v.Value}
	}
	m.EmployeeID = int16(req.GetEmployeeID())
	if v := req.GetExtension(); v != nil {
		m.Extension = sql.NullString{Valid: true, String: v.Value}
	}
	m.FirstName = req.GetFirstName()
	if v := req.GetHireDate(); v != nil {
		if err = v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid HireDate: %s%w", err.Error(), validation.ErrUserInput)
			return
		}
		if t := v.AsTime(); !t.IsZero() {
			m.HireDate.Valid = true
			m.HireDate.Time = t
		}
	}
	if v := req.GetHomePhone(); v != nil {
		m.HomePhone = sql.NullString{Valid: true, String: v.Value}
	}
	m.LastName = req.GetLastName()
	if v := req.GetNotes(); v != nil {
		m.Notes = sql.NullString{Valid: true, String: v.Value}
	}
	m.Photo = req.GetPhoto()
	if v := req.GetPhotoPath(); v != nil {
		m.PhotoPath = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetPostalCode(); v != nil {
		m.PostalCode = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetRegion(); v != nil {
		m.Region = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetReportsTo(); v != nil {
		m.ReportsTo = sql.NullInt64{Valid: true, Int64: v.Value}
	}
	if v := req.GetTitle(); v != nil {
		m.Title = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetTitleOfCourtesy(); v != nil {
		m.TitleOfCourtesy = sql.NullString{Valid: true, String: v.Value}
	}

	err = m.Upsert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}
