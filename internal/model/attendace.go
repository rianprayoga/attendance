package model

type AttendaceRequest struct {
	EmployeeId uint `json:"employeeId"`
}

type AttendaceResponse struct {
	EmployeeId  uint `json:"employeeId"`
	AttendaceId uint `json:"attendaceId"`
}

type TrxRecord struct {
	Id      uint `json:"id"`
	Amount  uint `json:"amount"`
	Success bool `json:"success"`
}
