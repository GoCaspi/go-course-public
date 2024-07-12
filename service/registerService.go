package service

import (
	"example-project/model"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . DatabaseInterface
type DatabaseInterface interface {
	InsertEmployees(employees []model.Employee) error
	GetEmployeeByID(id string) (*model.Employee, error)
}

type EmployeeService struct {
	DbService DatabaseInterface
}

func NewEmployeeService(dbInterface DatabaseInterface) *EmployeeService {
	return &EmployeeService{
		DbService: dbInterface,
	}
}

func (s *EmployeeService) CreateEmployees(employees []model.Employee) error {
	return s.DbService.InsertEmployees(employees)
}

func (s *EmployeeService) GetEmployeeById(id string) (*model.Employee, error) {
	return s.DbService.GetEmployeeByID(id)
}
