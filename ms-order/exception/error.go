package exception

import (
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorHandler struct {
	logger *zap.Logger
}

func NewErrorHandler(logger *zap.Logger) ErrorHandler {
	return ErrorHandler{
		logger: logger.WithOptions(zap.AddCallerSkip(1)),
	}
}

func (e ErrorHandler) logerror(err error) {
	e.logger.Error("server error", zap.Error(err))
}

func (e ErrorHandler) ErrInternal(err error) error {
	e.logerror(err)
	message := "the server encountered a problem and could not process your request"
	return status.Error(codes.Internal, message)
}

func (e ErrorHandler) ErrInvalidArgument(err error) error {
	return status.Error(codes.InvalidArgument, err.Error())
}

func (e ErrorHandler) ErrNotFound(customeMsg ...string) error {
	if len(customeMsg) == 0 {
		customeMsg = append(customeMsg, "the requested resource could not be found")
	}
	message := customeMsg[0]
	return status.Error(codes.NotFound, message)
}
