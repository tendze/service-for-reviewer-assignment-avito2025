package response

import (
	"encoding/json"
	"net/http"
)

const (
	TEAM_EXISTS  = "TEAM_EXISTS"
	PR_EXISTS    = "PR_EXISTS"
	PR_MERGED    = "PR_MERGED"
	NOT_ASSIGNED = "NOT_ASSIGNED"
	NO_CANDIDATE = "NO_CANDIDATE"
	NOT_FOUND    = "NOT_FOUND"

	BAD_REQUEST            = "BAD_REQUEST"
	INTERNAL_SERVER_ERRROR = "INTERNAL_SERVER_ERROR"
	SUCCESS                = "SUCCESS"
)

type ErrorObject struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error ErrorObject `json:"error"`
}

func JSONError(w http.ResponseWriter, statusCode int, code string, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error: ErrorObject{
			Code:    code,
			Message: message,
		},
	})
}

func JSONSuccess(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(data)
}
