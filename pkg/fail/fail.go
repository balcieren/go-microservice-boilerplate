package fail

import (
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Fail struct {
	Code    codes.Code `json:"code"`
	Message string     `json:"message"`
}

func New(code codes.Code, message ...string) error {
	var msg string
	if len(message) > 0 {
		msg = message[0]
	}

	return status.Error(codes.Code(code), msg)
}

func Convert(err error) (int, string) {
	if err == nil {
		return fiber.StatusOK, "OK"
	}
	st, ok := status.FromError(err)
	if !ok {
		return fiber.StatusInternalServerError, "Internal Server Error"
	}
	switch st.Code() {
	case codes.InvalidArgument:
		return fiber.StatusBadRequest, st.Message()
	case codes.NotFound:
		return fiber.StatusNotFound, st.Message()
	case codes.AlreadyExists:
		return fiber.StatusConflict, st.Message()
	case codes.PermissionDenied:
		return fiber.StatusForbidden, st.Message()
	case codes.Unauthenticated:
		return fiber.StatusUnauthorized, st.Message()
	case codes.Internal:
		return fiber.StatusInternalServerError, st.Message()
	case codes.Unimplemented:
		return fiber.StatusNotImplemented, st.Message()
	case codes.Unavailable:
		return fiber.StatusServiceUnavailable, st.Message()
	case codes.DeadlineExceeded:
		return fiber.StatusGatewayTimeout, st.Message()
	case codes.Canceled:
		return fiber.StatusRequestTimeout, st.Message()
	default:
		return fiber.StatusInternalServerError, st.Message()
	}
}
