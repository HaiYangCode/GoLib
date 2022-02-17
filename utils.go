package atlib

import (
	"encoding/json"
	"io"
	"net/http"
)

var result = make(map[string]interface{})

func RStatus(w http.ResponseWriter, statusCode int, statusMsg string, resultData interface{}) {
	result["code"] = statusCode
	result["msg"] = statusMsg
	result["data"] = resultData
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(result)
	io.WriteString(w, string(data))
}
