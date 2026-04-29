package model

type AttendaceRequest struct {
	EmployeeId uint `json:"employeeId"`
}

type AttendaceResponse struct {
	EmployeeId  uint `json:"employeeId"`
	AttendaceId uint `json:"attendaceId"`
}

type AttendaceHistoryResponse struct {
	AttendaceId uint   `json:"attendaceId"`
	CheckIn     string `json:"checkIn"`
	CheckOut    string `json:"checkOut"`
	Status      string `json:"status"`
}

type Page[T any] struct {
	Data []T `json:"data"`
}
