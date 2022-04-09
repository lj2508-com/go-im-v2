package util

import (
	"encoding/json"
	"net/http"
)

func RespFail(w http.ResponseWriter, msg string) {
	resp(w, -1, nil, msg)
}
func RespOk(w http.ResponseWriter, data interface{}, msg string) {
	resp(w, 0, data, msg)
}

func resp(writer http.ResponseWriter, code int, data interface{}, msg string) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	bady := respBady{
		Code: code,
		Data: data,
		Msg:  msg,
	}
	json, _ := json.Marshal(bady)
	writer.Write(json)
}

type respBady struct {
	Code int
	Data interface{}
	Msg  string
}
