package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/enescakir/emoji"
	"io"
	"net/http"
	pb "uDES/api/proto"
	"uDES/internal/logger"
	"uDES/internal/models"
)

const baseUrl = "http://localhost:8080/Portal/springApi/api/"

func doRequest(in interface{}, endpoint string) ([]byte, error) {
	if err := checkHealth(); err != nil {
		return nil, fmt.Errorf("http server is unavailable: %v", err)
	}
	jsonValue, err := json.Marshal(in)
	if err != nil {
		logger.WarErrLogger.Error().Err(err).Msg("Failed to do request to http server")
		return nil, err
	}

	req, err := http.NewRequest(
		"POST", baseUrl+endpoint, bytes.NewBuffer(jsonValue),
	)
	if err != nil {
		logger.WarErrLogger.Error().Err(err).Msg("Failed to do request to http server")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("admin", "qwerty")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		logger.WarErrLogger.Error().Err(err).Msg("Failed to do request to http server")
		return nil, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		logger.WarErrLogger.Error().Err(err).Msg("Failed to read response body")
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		switch response.StatusCode {
		case http.StatusUnauthorized:
			return nil, errors.New("unauthorized")
		case http.StatusNotFound:
			return nil, errors.New("not Found")
		case http.StatusInternalServerError:
			return nil, errors.New("internal Server Error")
		default:
			return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
		}
	}

	return data, nil
}

func checkHealth() error {
	response, err := http.Get("http://localhost:8080/health")
	if err != nil {
		return fmt.Errorf("failed to perform health check: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP server returned non-OK status: %d", response.StatusCode)
	}

	return nil
}

func HandleGetEmployee(in *pb.GetEmployeeParams) (*pb.EmployeesList, error) {
	var employeeResponse models.EmployeeResponse
	data, err := doRequest(in, "employees")
	if err != nil {
		logger.WarErrLogger.Error().Err(err).Msg("failed to get employee data")
		return nil, err
	}
	if err := json.Unmarshal(data, &employeeResponse); err != nil {
		logger.WarErrLogger.Error().Err(err).Msg("failed unmarshall json")
		return nil, err
	}

	var grpcEmployees []*pb.Employee
	for _, employeeData := range employeeResponse.Data {
		grpcEmployee := &pb.Employee{
			Id:          int32(employeeData.ID),
			DisplayName: employeeData.DisplayName,
			Email:       employeeData.Email,
			WorkPhone:   employeeData.WorkPhone,
		}
		grpcEmployees = append(grpcEmployees, grpcEmployee)
	}

	return &pb.EmployeesList{
		Data: grpcEmployees,
	}, nil
}

func HandleGetAbsences(in *pb.GetAbsencesParams) (*pb.AbsencesList, error) {
	var absenceResponse models.AbsencesResponse
	data, err := doRequest(in, "absences")
	if err != nil {
		logger.WarErrLogger.Error().Err(err).Msg("failed to get employee data")
		return nil, err
	}
	if err := json.Unmarshal(data, &absenceResponse); err != nil {
		logger.WarErrLogger.Error().Err(err).Msg("failed to unmarshal json")
		return nil, err
	}

	var grpcAbsences []*pb.Absence
	for _, absenceData := range absenceResponse.Data {
		grpcAbsence := &pb.Absence{
			Id:          int32(absenceData.ID),
			PersonId:    int32(absenceData.PersonID),
			CreatedDate: absenceData.CreatedDate,
			DateTo:      absenceData.DateTo,
			DateFrom:    absenceData.DateFrom,
			ReasonId:    int32(absenceData.ReasonID),
			DisplayName: getNameById(int32(absenceData.PersonID)) + getSmile(int(absenceData.ReasonID)),
		}
		grpcAbsences = append(grpcAbsences, grpcAbsence)
	}
	return &pb.AbsencesList{
		Data: grpcAbsences,
	}, nil
}

func getNameById(id int32) string {
	ids := []int32{id}
	EmployeeParams := &pb.GetEmployeeParams{
		Ids: ids,
	}
	employeeData, _ := HandleGetEmployee(EmployeeParams)
	var employeeName string
	for _, employee := range employeeData.Data {
		if employee.Id == id {
			employeeName = employee.DisplayName
			break
		}
	}
	return employeeName
}

func getSmile(reason int) string {
	switch reason {
	case 1, 10:
		return string(emoji.House)
	case 3, 4:
		return string(emoji.Airplane)
	case 5, 6:
		return string(emoji.Thermometer)
	case 9:
		return string(emoji.TopHat)
	case 11, 12, 13:
		return string(emoji.Sun)
	default:
		return ""
	}
}
