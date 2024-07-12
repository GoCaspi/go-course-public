package model

type CreateEmployeePayload struct {
	Employees []Employee `json:"employees" binding:"required,dive"`
}
