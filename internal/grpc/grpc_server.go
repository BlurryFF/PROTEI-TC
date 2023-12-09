package grpcserver

import (
	"context"
	"google.golang.org/grpc"
	"net"
	pb "uDES/api/proto"
	"uDES/internal/config"
	"uDES/internal/logger"
	"uDES/internal/workers"
)

func NewUserManagementServer() *UserManagementServer {
	cfg := config.LoadConfiguration("config.json")
	pool, err := workers.NewWorkerPool(cfg.GRPCConfig.NumWorkers, cfg.GRPCConfig.QueueSize)
	if err != nil {
		logger.WarErrLogger.Error().Err(err).Msg("Failed to create worker pool")
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
		logger.WarErrLogger.Error().Err(err).Msg("Failed to listen grpc server")
	}
	s := grpc.NewServer()
	server.WorkerPool.Start()
	pb.RegisterUserManagementServer(s, server)
	logger.GRPCLogger.Info().Msgf("gRPC server listening at ", server.Config.GRPCConfig.IP+server.Config.GRPCConfig.Port)
	return s.Serve(lis)
}

func (server *UserManagementServer) GetEmployee(_ context.Context, in *pb.GetEmployeeParams) (*pb.EmployeesList, error) {

	logger.GRPCLogger.Info().Any("gRPC request", in).Send()
	result := make(chan workers.TaskResult)

	req := &workers.GetTask{
		InParams: in,
		Result:   result,
	}

	server.WorkerPool.AddWork(req)

	taskResult := <-result

	employeeData, err := taskResult.GetData().(*pb.EmployeesList), taskResult.GetErr()
	if err != nil {
		logger.WarErrLogger.Error().Err(err).Msg("Failed to get employees data from http server")
		return nil, err
	}

	logger.GRPCLogger.Info().Any("gRPC response", employeeData).Send()

	return employeeData, nil
}

func (server *UserManagementServer) GetAbsences(_ context.Context, in *pb.GetAbsencesParams) (*pb.AbsencesList, error) {
	logger.GRPCLogger.Info().Any("gRPC request", in).Send()

	result := make(chan workers.TaskResult)

	req := &workers.GetTask{
		InParams: in,
		Result:   result,
	}

	server.WorkerPool.AddWork(req)

	taskResult := <-result
	absencesData, err := taskResult.GetData().(*pb.AbsencesList), taskResult.GetErr()
	if err != nil {
		logger.WarErrLogger.Error().Err(err).Msg("Failed to get absences data from http server")
		return nil, err
	}

	logger.GRPCLogger.Info().Any("gRPC response", absencesData).Send()

	return absencesData, nil
}
