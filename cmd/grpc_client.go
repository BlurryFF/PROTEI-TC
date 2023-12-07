package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
	pb "uDES/api/proto"
)

const (
	address = "127.0.0.1:8081"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	now := time.Now()
	params := &pb.GetEmployeeParams{
		Ids:       []int32{2, 3},
		Email:     "",
		Name:      "",
		WorkPhone: "",
		DateFrom:  time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).Format("2006-01-02T15:04:05Z"),
		DateTo:    time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.UTC).Format("2006-01-02T15:04:05Z"),
	}

	paramsAbs := &pb.GetAbsencesParams{
		PersonIds: []int32{},
		DateFrom:  time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02T15:04:05Z"),
		DateTo:    time.Date(2023, 12, 5, 0, 0, 0, 0, time.UTC).Format("2006-01-02T15:04:05Z"),
	}

	absenceRes, err := c.GetAbsences(ctx, paramsAbs)
	if err != nil {
		log.Fatalf("could not get absences: %v", err)
	}
	log.Println("Abcences: ", absenceRes.GetData())

	employeeRes, err := c.GetEmployee(ctx, params)
	if err != nil {
		log.Fatalf("could not get employees: %v", err)
	}
	log.Println("Employees: ", employeeRes.GetData())
}
