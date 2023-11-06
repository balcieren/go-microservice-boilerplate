package utils

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ConvertGRPCErrorToHTTP(err error) (int, string) {
	if err == nil {
		return http.StatusOK, "OK"
	}
	st, ok := status.FromError(err)
	if !ok {
		return http.StatusInternalServerError, "Internal Server Error"
	}
	switch st.Code() {
	case codes.InvalidArgument:
		return http.StatusBadRequest, st.Message()
	case codes.NotFound:
		return http.StatusNotFound, st.Message()
	case codes.AlreadyExists:
		return http.StatusConflict, st.Message()
	case codes.PermissionDenied:
		return http.StatusForbidden, st.Message()
	case codes.Unauthenticated:
		return http.StatusUnauthorized, st.Message()
	case codes.Internal:
		return http.StatusInternalServerError, st.Message()
	case codes.Unimplemented:
		return http.StatusNotImplemented, st.Message()
	case codes.Unavailable:
		return http.StatusServiceUnavailable, st.Message()
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout, st.Message()
	case codes.Canceled:
		return http.StatusRequestTimeout, st.Message()
	default:
		return http.StatusInternalServerError, st.Message()
	}
}
