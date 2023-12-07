package models

type Employee struct {
	ID          int    `json:"id"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	MobilePhone string `json:"mobilePhone"`
	WorkPhone   string `json:"workPhone"`
}

type Absences struct {
	ID          int    `json:"id"`
	PersonID    int    `json:"personId"`
	CreatedDate string `json:"createdDate"`
	DateFrom    string `json:"dateFrom"`
	DateTo      string `json:"dateTo"`
	ReasonID    uint8  `json:"reasonId"`
}

type EmployeeRequest struct {
	Ids       []int  `json:"ids"`
	Name      string `json:"name"`
	WorkPhone string `json:"workPhone"`
	Email     string `json:"email"`
	DateFrom  string `json:"dateFrom"`
	DateTo    string `json:"dateTo"`
}

type EmployeeNameRequest struct {
	PersondIds []int `json:"persondIds"`
}

type EmployeeNameResponse struct {
	displayNames map[int]string `json:"displayNames"`
}

type EmployeeResponse struct {
	Status string     `json:"status"`
	Data   []Employee `json:"data"`
}

type AbsencesRequest struct {
	PersonIds []int  `json:"personIds"`
	DateFrom  string `json:"dateFrom"`
	DateTo    string `json:"dateTo"`
}

type AbsencesResponse struct {
	Status string     `json:"status"`
	Data   []Absences `json:"data"`
}
