package models

import (
	"encoding/json"
	"net/http"
)

const (
	dbExecContext = "Error exec context"
)

var (
	DbExecContext = getError(dbExecContext, -15)
)

type ErrorResponse struct {
	ErrorCode     int    `json:"errorCode"`
	ErrorDesc     string `json:"errorDescription"`
	AditionalInfo Imodel `json:"response,omitempty"`
}

func (r ErrorResponse) Error() string {
	return r.toJson()
}

func (r ErrorResponse) toJson() string {
	reply, err := json.Marshal(r)
	if err != nil {
		return "{error:\"Marshal error\"}"
	}
	return string(reply)
}

func ToJson(w http.ResponseWriter, status int, data interface{}) {

	reply, err := json.Marshal(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(reply)
}

func getError(err string, code int) ErrorResponse {
	return ErrorResponse{
		ErrorCode: code,
		ErrorDesc: err,
	}
}

func SetError(err error) ErrorResponse {
	return ErrorResponse{
		ErrorDesc: err.Error(),
	}
}
