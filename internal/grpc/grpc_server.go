package grpcserver

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "uDES/api/proto"
	"uDES/internal/config"
	"uDES/internal/workers"
)

func NewUserManagementServer() *UserManagementServer {
	cfg := config.LoadConfiguration("config.json")
	pool, err := workers.NewWorkerPool(cfg.GRPCConfig.NumWorkers, cfg.GRPCConfig.QueueSize)
	if err != nil {
		log.Printf("failed to create worker pool: %v", err)
	}

	return &UserManagementServer{
		WorkerPool: pool,
		Config:     cfg,
	}
}

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
	WorkerPool workers.Pool
	Config     config.Config
}

func (server *UserManagementServer) Run() error {
	lis, err := net.Listen("tcp", server.Config.GRPCConfig.IP+server.Config.GRPCConfig.Port)
	if err != nil {
		log.Printf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	server.WorkerPool.Start()
	pb.RegisterUserManagementServer(s, server)
	log.Printf("gRPC server listening at %v", lis.Addr())
	return s.Serve(lis)
}

func (server *UserManagementServer) GetEmployee(_ context.Context, in *pb.GetEmployeeParams) (*pb.EmployeesList, error) {
	result := make(chan workers.TaskResult)

	req := &workers.GetTask{
		InParams: in,
		Result:   result,
	}

	server.WorkerPool.AddWork(req)

	taskResult := <-result

	employeeData, err := taskResult.GetData().(*pb.EmployeesList), taskResult.GetErr()
	if err != nil {
		return nil, err
	}

	return employeeData, nil
}

func (server *UserManagementServer) GetAbsences(_ context.Context, in *pb.GetAbsencesParams) (*pb.AbsencesList, error) {
	result := make(chan workers.TaskResult)

	req := &workers.GetTask{
		InParams: in,
		Result:   result,
	}

	server.WorkerPool.AddWork(req)

	taskResult := <-result
	absencesData, err := taskResult.GetData().(*pb.AbsencesList), taskResult.GetErr()
	if err != nil {
		return nil, err
	}
	return absencesData, nil
}
