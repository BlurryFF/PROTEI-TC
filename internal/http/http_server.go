package httpserver

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"uDES/internal/config"
	"uDES/internal/models"
)

type HTTPServer struct {
	server *http.Server
}

var cfg *config.Config

func healthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		return
	}
}

func employeeInfoHandler(w http.ResponseWriter, r *http.Request) {
	if !verifyUser(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	var request models.EmployeeRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	employees := employeeFinder(request)

	if len(employees) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response := struct {
		Status string            `json:"status"`
		Data   []models.Employee `json:"data"`
	}{
		Status: "OK",
		Data:   employees,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func absencesInfoHandler(w http.ResponseWriter, r *http.Request) {
	if !verifyUser(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	var request models.AbsencesRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	absences := absenceFinder(request)
	if len(absences) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	response := struct {
		Status string            `json:"status"`
		Data   []models.Absences `json:"data"`
	}{
		Status: "OK",
		Data:   absences,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func loadAbsencesData() ([]models.Absences, error) {
	absencesFilePath := filepath.Join(cfg.HTTPConfig.BaseStoragePath, cfg.HTTPConfig.AbsencesFileName)
	data, err := os.ReadFile(absencesFilePath)
	if err != nil {
		return nil, err
	}

	var absences []models.Absences
	err = json.Unmarshal(data, &absences)
	if err != nil {
		return nil, err
	}

	return absences, nil
}

func loadEmployeesData() ([]models.Employee, error) {
	absencesFilePath := filepath.Join(cfg.HTTPConfig.BaseStoragePath, cfg.HTTPConfig.EmployeesFileName)
	data, err := os.ReadFile(absencesFilePath)
	if err != nil {
		return nil, err
	}

	var employees []models.Employee
	err = json.Unmarshal(data, &employees)
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func absenceFinder(request models.AbsencesRequest) (absences []models.Absences) {
	Absences, _ := loadAbsencesData()
	layout := "2006-01-02T15:04:05Z"
	reqDateFrom, _ := time.Parse(layout, request.DateFrom)
	reqDateTo, _ := time.Parse(layout, request.DateTo)
	for _, absence := range Absences {
		absDateFrom, _ := time.Parse(layout, absence.DateFrom)
		absDateTo, _ := time.Parse(layout, absence.DateTo)
		if absDateFrom.After(reqDateFrom) && absDateTo.Before(reqDateTo) {
			if len(request.PersonIds) > 0 {
				for _, id := range request.PersonIds {
					if id == absence.PersonID {
						absences = append(absences, absence)
					}
				}
			} else {
				absences = append(absences, absence)
			}
		}
	}
	return absences
}

func employeeFinder(request models.EmployeeRequest) (employees []models.Employee) {
	Employees, _ := loadEmployeesData()
	if len(request.Ids) == 0 && request.Name == "" && request.WorkPhone == "" && request.Email == "" {
		return employees
	}

	if len(request.Ids) > 0 {
		for _, id := range request.Ids {
			for _, employee := range Employees {
				if employee.ID == id {
					employees = append(employees, employee)
				}
			}
		}
		if len(employees) > 0 {
			return employees
		}
	}

	if request.Email != "" {
		for _, employee := range Employees {
			if employee.Email == request.Email {
				employees = append(employees, employee)
				if len(employees) > 0 {
					return employees
				}
			}
		}
	}

	if request.Name != "" {
		for _, employee := range Employees {
			if employee.DisplayName == request.Name {
				employees = append(employees, employee)
				if len(employees) > 0 {
					return employees
				}
			}
		}
	}

	if request.WorkPhone != "" {
		for _, employee := range Employees {
			if employee.WorkPhone == request.WorkPhone {
				employees = append(employees, employee)
				if len(employees) > 0 {
					return employees
				}
			}
		}
	}

	return employees
}

func verifyUser(r *http.Request) bool {
	u, p, ok := r.BasicAuth()
	if u == cfg.HTTPConfig.Auth.Username && p == cfg.HTTPConfig.Auth.Pass && ok {
		return true
	}
	return false
}

func NewHTTPServer() *HTTPServer {
	cfgLoad := config.LoadConfiguration("config.json")
	cfg = &cfgLoad
	mux := http.NewServeMux()

	mux.HandleFunc("/Portal/springApi/api/employees", employeeInfoHandler)
	mux.HandleFunc("/Portal/springApi/api/absences", absencesInfoHandler)
	mux.HandleFunc("/health", healthCheckHandler)

	server := http.Server{
		Addr:    cfg.HTTPConfig.IP + cfg.HTTPConfig.Port,
		Handler: mux,
	}

	return &HTTPServer{server: &server}
}

func (hs *HTTPServer) Start() error {
	log.Printf("HTTP server listening at %v", hs.server.Addr)
	return hs.server.ListenAndServe()
}

func (hs *HTTPServer) Stop() error {
	log.Println("Stopping HTTP server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return hs.server.Shutdown(ctx)
}
