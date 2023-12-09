package logger

import (
	"github.com/rs/zerolog"
	"os"
)

var (
	logger       zerolog.Logger
	GRPCLogger   zerolog.Logger
	HTTPLogger   zerolog.Logger
	WarErrLogger zerolog.Logger
	DebugLogger  zerolog.Logger
)

func init() {
	logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

	grpcFile, err := os.OpenFile("logs/grpc_server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to create/open gRPC log file")
	}
	GRPCLogger = logger.Output(grpcFile).Level(zerolog.InfoLevel)

	httpFile, err := os.OpenFile("logs/http_server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to create/open http log file")
	}
	HTTPLogger = logger.Output(httpFile).Level(zerolog.InfoLevel)

	warErrFile, err := os.OpenFile("logs/warning_error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to create/open warning/error file")
	}
	WarErrLogger = logger.Output(warErrFile)

	debugFile, err := os.OpenFile("logs/debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to create/open debug file")
	}
	DebugLogger = logger.Output(debugFile).Level(zerolog.DebugLevel)
}
